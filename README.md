

# ① 先把 README.md 重新规整

我建议统一成**中文为主、代码和专业术语保留英文**。

直接替换你的 `README.md`：

# ged

`ged` 是一个基于 Go 标准库构建的轻量级 Edge Runtime Core。

它用于帮助 Go 应用构建：

```text
Consumer
    ↓
HTTP Edge
    ↓
Adapter / BFF
    ↓
gRPC
    ↓
Provider
```

ged 不负责业务逻辑，而是提供 Edge 层中可以复用的基础运行能力。

------

## 定位

ged 的核心目标是：

> 保持 Edge 足够薄、明确、类型安全，并尽可能使用 Go 标准库。

ged 主要关注以下能力：

- Request Context
- Identity Context
- Authentication Extension Point
- gRPC Metadata Propagation
- gRPC Client
- 基础并发编排

业务 API、业务规则和 Provider 实现不属于 ged。

------

## 核心结构

```
                    Consumer
                       │
                       ▼
                    net/http
                       │
              ┌────────┴────────┐
              │                 │
       Authentication       Request Context
              │                 │
              ▼                 ▼
          Identity          Request ID
              │              Trace ID
              │                 │
              └────────┬────────┘
                       │
                       ▼
                 HTTP Adapter
                       │
                       ▼
                  gRPC Client
                       │
              Metadata Propagation
                       │
                       ▼
                   Provider
```

对于需要聚合多个 Provider 的 API：

```
                    Consumer
                       │
                       ▼
                      BFF
                       │
              ┌────────┴────────┐
              ▼                 ▼
          Provider A        Provider B
              │                 │
              └────────┬────────┘
                       ▼
                   Response
```

------

## 核心模块

### authentication

定义认证能力的扩展点。

ged 不实现具体的认证机制。

```go
type Authenticator interface {
    Authenticate(token string) (identity.Context, error)
}
```

应用程序负责提供具体实现。

------

### identity

提供当前请求的身份上下文。

目前包含：

```
UserID
TenantID
```

身份信息通过 `context.Context` 在 Edge 内部传播，并可以进一步传播到 gRPC Metadata。

------

### request

提供当前 HTTP Request 的上下文信息。

目前包含：

```
RequestID
TraceID
```

支持：

```
HTTP Header
    ↓
context.Context
    ↓
gRPC Metadata
```

------

### grpc

提供 Edge → Provider 的 gRPC Client 能力。

主要负责：

- 创建 gRPC Client Connection
- 自动传播 Request Context
- 自动传播 Identity
- 将 Edge Context 转换为 gRPC Metadata

应用层不需要关心 Metadata 的具体传播过程。

------

### orchestration

提供轻量级的并发编排能力。

例如：

```go
orchestration.ParallelFailFast(...)
```

它只负责并发执行与取消控制。

具体的业务编排仍然属于应用层。

------

## HTTP Edge

ged 使用 Go 标准库：

```go
net/http
```

而不是绑定某个 HTTP Framework。

Edge Adapter 通常负责：

```
HTTP Request
    ↓
Request Binding
    ↓
HTTP → gRPC Request
    ↓
gRPC Call
    ↓
gRPC Response
    ↓
HTTP Response
```

例如：

```go
func (a *Adapter) sayHello(
    w http.ResponseWriter,
    r *http.Request,
) {
    // HTTP → gRPC
}
```

这种 HTTP Adapter 属于具体应用，而不是 ged Runtime。

------

## Adapter

Adapter 的职责是协议转换。

例如：

```
HTTP
 ↓
Adapter
 ↓
gRPC
```

Adapter 不应该承担业务逻辑。

它主要负责：

- HTTP Request Binding
- HTTP → gRPC Request
- 调用 Provider
- gRPC Response → HTTP Response
- HTTP 特有的协议处理，例如 SSE

------

## BFF / Aggregate

当一个 Consumer API 需要调用多个 Provider 时，可以在 Edge 层实现 BFF / Aggregate。

例如：

```
GET /api/v1/dashboard/overview
            │
            ▼
        Dashboard
         Service
        /       \
       ▼         ▼
 DictArea      Hello
 Provider      Provider
       \         /
        ▼       ▼
          Response
```

BFF 属于应用层。

ged 只提供必要的运行时能力，例如：

```go
orchestration.ParallelFailFast(...)
```

------

## Go 标准库

ged 尽可能使用 Go 标准库。

当前 HTTP Edge 使用：

```go
net/http
http.ServeMux
http.Handler
http.ResponseWriter
```

因此 ged 不依赖 Gin、Echo 等 HTTP Framework。

------

## 示例

`examples/edge` 提供一个完整的 Edge 示例。

包含：

```
examples/edge
├── direct
│   ├── hello
│   └── dictarea
│
├── aggregate
│   └── dashboard
│
├── clients
├── authentication
├── logging
└── config
```

示例用于展示 ged 的实际使用方式。

它们不属于 ged Runtime API。

------

## Go 版本

当前项目使用：

```
Go 1.27
```

Go 新语言特性会在：

```
experiments/
```

中进行独立实验。

实验代码不会因为语言特性存在，就自动进入 ged 公共 API。

------

## 开发

运行测试：

```
go test ./...
```

运行静态检查：

```
go vet ./...
```

运行 Edge 示例：

```
go run ./examples/edge/
```

------

## 项目原则

ged 遵循几个简单原则：

1. 优先 Go 标准库
2. 保持 Edge 薄
3. Adapter 只做协议转换
4. 业务逻辑留在应用层
5. 不为了复用而过度抽象
6. 不因为语言支持某个特性就强行使用
7. 先解决真实重复，再决定是否抽象

这个版本我认为比之前合适很多。

---

# ② `v0.1 Frozen Scope` 不放 README

是的。

我们以后可以单独建立：

```text
docs/
└── adr/
    └── ADR-0001-ged-v0.1-boundary.md
```
