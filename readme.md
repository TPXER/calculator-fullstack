# 全栈开发挑战：计算器（Go + ConnectRPC + Next.js）

## 项目说明

本项目实现了一个支持加减乘除的全栈计算器，使用 Go 编写后端服务，Next.js 构建前端页面，并通过 ConnectRPC 实现前后端通信。

---

## 技术栈

- 后端：Go + ConnectRPC + Protobuf
- 前端：Next.js（React）
- 协议：Connect（非传统 REST）

---

## 主要功能

- 加减乘除基本运算
- 前端页面输入、选择运算符并展示结果
- 前后端通过 ConnectRPC 协议交互

---

## 运行说明

1. 安装依赖（需要预装以下环境）：
   - Go 1.20+
   - Node.js 18+
   - `protoc`（Protocol Buffers 编译器）
   - `protoc-gen-go` 与 `protoc-gen-connect-go`

2. 启动后端：

   ```bash
   cd backend
   go run main.go

启动前端：
cd frontend
npm install
npm run dev

文件结构
fullstack_calculator/
├── backend/             # Go 服务端
│   ├── main.go          # 服务入口
│   ├── calculator.proto # Protobuf 接口定义
│   └── go.mod / go.sum
├── frontend/            # Next.js 前端界面
│   └── src/app/page.tsx
└── readme.md            # 当前说明文档
