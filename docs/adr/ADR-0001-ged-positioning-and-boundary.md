很好。现在这个状态我建议**正式停止继续“搬家”**，进入 `ged v0.1 Freeze`。

这一版的目标不是“功能更多”，而是：

> **把 `ged` 从 goro-edge 的内部代码，冻结成一个边界清晰、可测试、可解释的 Go Library。** 

我建议这一轮只做：

```text
ged v0.1
├── ① API Freeze
├── ② 补最小单元测试
├── ③ README
└── ④ ADR-0001：定位与边界
```

暂时**不做**：

```text
❌ Error Boundary
❌ OpenTelemetry
❌ Rate Limit
❌ Retry
❌ Circuit Breaker
❌ Config
❌ Router
❌ BFF abstraction
❌ Gateway abstraction
```

这些全部留给后面。

------

# 一、先冻结 `ged v0.1` 的目录

我建议最终保持：

```text
ged/
├── go.mod
├── README.md
│
├── authentication/
│   ├── authenticator.go
│   ├── middleware.go
│   └── middleware_test.go
│
├── grpc/
│   ├── client.go
│   ├── metadata.go
│   └── metadata_test.go
│
├── http/
│   ├── handler.go
│   └── handler_test.go
│
├── identity/
│   ├── context.go
│   └── context_test.go
│
└── request/
    ├── context.go
    ├── metadata.go
    ├── middleware.go
    └── context_test.go
```

注意：

> **没有 `observability/`。**

这是我们刚刚已经确定的边界。

------

# 二、第一件事：补 `request.Middleware`

现在 `request` 里面只有：

```text
request/
├── context.go
└── metadata.go
```

但是 `RequestID / TraceID` 从 HTTP 进入 Context 的动作，本身就是 `ged` 的基础能力。

所以我们补：`ged/request/middleware.go` 

这样 `ged/request` 就形成完整闭环：

```text
HTTP
 │
 │ X-Request-ID
 │ X-Trace-ID
 ▼
request.Middleware
 │
 ▼
request.Context
 │
 ▼
gRPC Metadata
```

------

# 三、现在给 `identity` 补测试

`ged/identity/context_test.go` 这个测试非常简单。

但它在告诉我们一个非常重要的事实：

> `identity.Context` 是 `ged` 的 Context Primitive。

------

# 四、给 `request` 补测试

`ged/request/context_test.go` 

------

# 五、测试 Request → gRPC Metadata

这个测试其实是 `ged v0.1` 最重要的测试之一。

`ged/grpc/metadata_test.go` 

------

# 六、测试 Authentication

这里我们要区分：`Authenticator` 和 `Middleware` 

`ged/authentication/middleware_test.go` 

这里我们没有测试：DemoAuthenticator，因为它只是 demo implementation。

真正需要冻结的是：

```text
Authenticator interface
Middleware behavior
```

------

# 七、测试 JSONHandler

这个也是 `ged/http` 的核心 API。

`ged/http/handler_test.go` 

------

# 八、这里我建议您暂时不要测试 `grpc.NewClient`

原因很简单：

```go
NewClient()
```

属于：

```text
transport wiring
```

真正需要测试的是：

```text
PropagateMetadata()
```

而不是测试：

> “gRPC 官方 ClientConn 能不能创建。”

这不是 `ged` 的责任。

这体现一个原则：

> **只测试自己的行为，不测试依赖库已经保证的行为。**

------

# 九、现在运行测试

在：

```text
D:\roc-github\goro-edge
```

执行：

```powershell
go test ./...
```

然后：

```powershell
go test ./ged/...
```

如果两个都通过：

```text
PASS
```

我们就可以认为：

```text
ged v0.1
```

开始具有最小的工程可信度。

------

# 十、接下来写 README

## `../../example/edge.md`

我建议第一版不要写长。

这份 README 有一个很重要的作用：

> **防止未来的我们把 `ged` 越做越胖。**

---

# 十一、然后写 ADR-0001

我认为这个 ADR 比代码更重要。

建议：

```ini
goro-edge/
└── ged/
    └── docs/
        └── adr/
            └── ADR-0001-ged-positioning-and-boundary.md
```

完整内容：

# ADR-0001: ged 定位与边界

- Status: Accepted
- Date: 2026-08-18

## Context

`goro-edge` 是一个具体的 Edge Application。

在实际开发过程中发现，不同 Edge Application 会重复出现一组与具体业务无关的基础能力：

- Request ID / Trace ID
- Identity Context
- HTTP Request Context
- gRPC Metadata Propagation
- Authentication Middleware Boundary
- 通用 HTTP JSON Handler

这些能力既不属于具体业务 Application，也不应该散落在每个项目中。

因此需要将其抽取为一个独立的 Go module。

## Decision

创建独立 Go module：

```markdown
github.com/rocwg/ged
```

项目名称为：

```text
ged
```

其定位为：

> 一个面向 Go Edge Application 的轻量、类型安全基础库。

## Responsibilities

`ged` 负责提供 Edge Application 的基础 primitives：

### Request

负责：

- Request ID
- Trace ID
- Request Context
- Request Context → gRPC Metadata

### Identity

负责：

- Identity Context
- Identity Context → gRPC Metadata

### gRPC

负责：

- Edge → Provider gRPC Client
- Outgoing Metadata Propagation

### HTTP

负责：

- 通用 HTTP JSON Handler

### Authentication

负责：

- Authentication interface
- Authentication Middleware

认证实现由 Application 提供。

## Non-Responsibilities

`ged` 不负责：

- Business Logic
- API Route Definition
- BFF
- Provider Client Registry
- Configuration
- Logging Policy
- Metrics Policy
- Tracing Backend
- Error Policy
- Rate Limiting
- Retry
- Circuit Breaking
- Load Balancing
- Reverse Proxy
- Service Discovery
- Service Mesh

## Application Ownership

具体 Edge Application 保留以下控制权：

```text
Application
├── Routes
├── BFF
├── Aggregate
├── Direct Adapter
├── Provider Clients
├── Configuration
├── Authentication Implementation
├── Logging Policy
├── Observability Policy
└── Error Policy
```

`ged` 只提供基础能力。

## Design Principles

### Small

优先保持 API 数量少。

如果某项能力只在单个 Application 中使用，默认不进入 `ged`。

### Type-safe

优先使用：

- Go types
- interfaces
- generics

避免通过：

- map[string]any
- reflection
- dynamic configuration

构建核心 API。

### Go-native

遵循 Go 的基本设计：

- simple APIs
- explicit dependencies
- context.Context
- interfaces at boundaries
- composition over framework inheritance

### Business-unaware

`ged` 不应该知道：

- IAM
- CMS
- Order
- DictArea
- Dashboard

等业务概念。

## Relationship with goro-edge

`goro-edge` 是具体 Application。

```text
goro-edge
    ├── applications
    └── ged
```

其中：

```text
applications
    ↓
使用 ged
```

`ged` 不反向依赖：

```text
goro-edge/applications
```

依赖关系必须保持单向：

```text
Application
     │
     ▼
    ged
```

## Relationship with KrakenD / Gateway Products

`ged` 不以重新实现 KrakenD、Envoy、Kong 等 Gateway 产品为目标。

如果未来需要：

- advanced routing
- rate limiting
- load balancing
- circuit breaking
- proxying
- service discovery

应优先考虑成熟基础设施，而不是持续扩大 `ged` 的职责。

## Versioning

`ged v0.1` API 保持最小化。

在没有明确使用场景之前，不增加新的抽象。

新能力必须满足：

1. 在实际 Edge Application 中出现真实重复需求；
2. 与 Edge 基础职责直接相关；
3. 可以保持简单；
4. 不引入不必要的业务耦合。

## Consequence

### Positive

- Edge Application 更轻
- 基础能力可复用
- HTTP → Context → gRPC Metadata 链路统一
- API 边界更加清晰
- Application 保留业务控制权

### Negative

- 需要维护独立 module
- 需要考虑 API compatibility
- 过度抽象会增加维护成本

因此 `ged` 必须优先保持小，而不是追求功能完整。

---

# 十二、这里其实完成了一个非常重要的架构转变

以前是：

```ini
goro-edge
├── runtime
└── extensions
```

我们容易产生这样的心理：

> “这些是我的 Edge 内部基础设施。”

现在变成：

```text
goro-edge
├── applications
└── ged
```

语义已经完全不同：

```text
goro-edge
   └── 一个使用 ged 的真实 Edge Application
```

而 `ged` 本身：

```text
ged
│
├── request
├── identity
├── grpc
├── http
└── authentication
```

是：

> **一个独立的 Go library。**

这就是您之前说的：

> “不是重新造一个 KrakenD，而是做一个足够薄、类型安全、自己真正掌控的专用于 Edge 层的封装库。”

**现在我们已经真正开始把这个想法落地了。**

------

# 十三、这一轮最后做一个 Freeze Check

全部完成后，您执行：

```powershell
cd D:\roc-github\goro-edge

go work sync
go test ./...
go vet ./...
```

然后：

```powershell
git status
```

再看：

```powershell
git diff
```

最后建议形成一个非常干净的 commit：

```text
feat: freeze ged v0.1
```

然后打 tag：

```powershell
git tag ged-v0.1.0
```

不过**先不要急着 tag**。

我们下一步应该先看您实际写出来的 `ged` 当前代码，尤其是 `request.Middleware`、`authentication`、`grpc` 和 `http.JSONHandler` 四个 API。

**我会在 Freeze 前帮您做一次“API 审查”，确认哪些函数名、类型名、参数形式值得现在改掉。**

因为：

> **v0.1 之前改 API 是重构；v0.1 之后改 API 就开始产生兼容性成本。** 

这也是现在最值得花时间的一步。

