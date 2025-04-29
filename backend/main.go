package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "calculator/gen"
    "calculator/genconnect"

    "connectrpc.com/connect"
    "golang.org/x/net/http2"
    "golang.org/x/net/http2/h2c"
)

// 实现 CalculatorServiceHandler 接口
type CalculatorServer struct{}

func (s *CalculatorServer) Add(ctx context.Context, req *connect.Request[gen.CalculatorRequest]) (*connect.Response[gen.CalculatorResponse], error) {
    a := req.Msg.A
    b := req.Msg.B
    result := a + b
    return connect.NewResponse(&gen.CalculatorResponse{Result: result}), nil
}

func (s *CalculatorServer) Subtract(ctx context.Context, req *connect.Request[gen.CalculatorRequest]) (*connect.Response[gen.CalculatorResponse], error) {
    a := req.Msg.A
    b := req.Msg.B
    result := a - b
    return connect.NewResponse(&gen.CalculatorResponse{Result: result}), nil
}

func (s *CalculatorServer) Multiply(ctx context.Context, req *connect.Request[gen.CalculatorRequest]) (*connect.Response[gen.CalculatorResponse], error) {
    a := req.Msg.A
    b := req.Msg.B
    result := a * b
    return connect.NewResponse(&gen.CalculatorResponse{Result: result}), nil
}

func (s *CalculatorServer) Divide(ctx context.Context, req *connect.Request[gen.CalculatorRequest]) (*connect.Response[gen.CalculatorResponse], error) {
    a := req.Msg.A
    b := req.Msg.B
    if b == 0 {
        return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("cannot divide by zero"))
    }
    result := a / b
    return connect.NewResponse(&gen.CalculatorResponse{Result: result}), nil
}

func main() {
    mux := http.NewServeMux()
    path, handler := genconnect.NewCalculatorServiceHandler(&CalculatorServer{})
    mux.Handle(path, handler)

    // ✅ 添加 CORS 支持
    corsWrapper := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        mux.ServeHTTP(w, r)
    })

    fmt.Println("✅ Calculator RPC server is running on http://localhost:8080")

    // 使用 h2c 包装 CORS handler
    if err := http.ListenAndServe(
        "localhost:8080",
        h2c.NewHandler(corsWrapper, &http2.Server{}),
    ); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
