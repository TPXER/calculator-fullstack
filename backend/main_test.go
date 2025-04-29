package main

import (
	"context"
	"testing"

	"calculator/gen"

	"connectrpc.com/connect"
)

func TestAdd(t *testing.T) {
	server := &CalculatorServer{}
	req := connect.NewRequest(&gen.CalculatorRequest{A: 10, B: 5})
	res, err := server.Add(context.Background(), req)

	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	if res.Msg.Result != 15 {
		t.Errorf("Expected 15, got %v", res.Msg.Result)
	}
}

func TestSubtract(t *testing.T) {
	server := &CalculatorServer{}
	req := connect.NewRequest(&gen.CalculatorRequest{A: 10, B: 5})
	res, err := server.Subtract(context.Background(), req)

	if err != nil {
		t.Fatalf("Subtract failed: %v", err)
	}
	if res.Msg.Result != 5 {
		t.Errorf("Expected 5, got %v", res.Msg.Result)
	}
}

// Multiply 和 Divide 同样编写即可
