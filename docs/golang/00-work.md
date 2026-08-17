很好，我们正式开始。

不过，我想先**调整一个细节**，这样以后不会踩坑。

## 我建议先不要创建 `shared/`

原因是：

Go 官方团队和很多大型 Go 项目（如 Kubernetes、Kratos 等）都有一个共识：

> **不要预先设计共享目录（shared/common/utils），只有当真正出现重复代码时，再抽取。**

为什么？

因为很多项目一开始就建：

```text
shared/
common/
util/
base/
```

半年以后：

```text
shared/
    config/
    dto/
    grpcclient/
    util/
    helper/
    common/
```

最后变成"垃圾场"（Dumping Ground）。

所以，我建议：**V1 不创建 `shared/`。** 

等 Adapter 和 BFF 都真正需要同一份代码时，我们再抽取。

------

# 我建议的 V1 目录

```text
goro-edge/
│
├── go.work
│
├── apps/
│   ├── adapter/
│   │   ├── go.mod
│   │   ├── cmd/
│   │   │   └── adapter/
│   │   │       └── main.go
│   │   ├── internal/
│   │   └── README.md
│   │
│   └── bff/
│       ├── go.mod
│       ├── cmd/
│       │   └── bff/
│       │       └── main.go
│       ├── internal/
│       └── README.md
│
├── docs/
│
├── .gitignore
└── README.md
```

注意：

目前：**没有 shared。** 

------

# 为什么？

因为：

现在：Adapter 和 BFF

实际上：没有一行代码。

共享什么？

不知道。

等以后：

例如：

都需要：JWT Middleware

再：internal $\rightarrow$ pkg $\rightarrow$ shared 都可以。

Go 一直提倡：

> **抽象来自重复，而不是来自预测。**

这是我非常认同的一句话。

------

# Phase 1：创建目录

仓库：goro-edge

```text
goro-edge
├── apps
│   ├── adapter
│   └── bff
├── docs
└── README.md
```

先建立即可。不用写代码。

README：先一句话即可：

```markdown
# goro-edge

Edge Layer for Pure gRPC Services.

Contains:

- HTTP Adapter
- Backend For Frontend (BFF)
```

------



# Phase 2：git 初始化

```bash
git init
git add .
git commit -m "init repository"
```

------



# Phase 3：初始化 Adapter

进入：apps/adapter

执行：go mod init github.com/rocwg/goro-edge/apps/adapter

然后：cmd/adapter/main.go：

```go
package main

import "fmt"

func main() {
	fmt.Println("goro-edge adapter starting...")
}
```

结束。不要多写

------



# Phase 4：初始化 BFF

同样：apps/bff

执行：go mod init github.com/rocwg/goro-edge/apps/bff

然后：cmd/bff/main.go：

```go
package main

import "fmt"

func main() {
	fmt.Println("goro-edge bff starting...")
}
```

结束。

------

# Phase 5：创建 go.work

回到：goro-edge/

执行：go work init

然后：

```bash
go work use ./apps/adapter
go work use ./apps/bff
```

最终：go.work

应该类似：

```go
go 1.25

use (
    ./apps/adapter
    ./apps/bff
)
```

------



# Phase 6：验证

根目录：执行：go work sync

然后：

Adapter：

```bash
cd apps/adapter
go run ./cmd/adapter
```

输出：

```
goro-edge adapter starting...
```

------

BFF：

```bash
cd ../bff
go run ./cmd/bff
```

输出：

```
goro-edge bff starting...
```

说明：整个 Workspace 已经正常。

------



# 为什么我推荐两个 Module？

因为：以后：Docker：

Adapter：

```
FROM golang

COPY apps/adapter
```

BFF：

```
FROM golang

COPY apps/bff
```

完全独立。

------

CI：

可以：

只编译：

```
adapter
```

或者：

```
bff
```

互不影响。

------



# 什么时候出现 shared？

例如：

半年以后：

发现：

Adapter：有：JWT

BFF：也有：JWT

再：

抽：pkg/auth

或者：shared/auth

都可以。

这是：**基于事实抽象（Fact-based Abstraction）**。

不是：**预测抽象（Prediction-based Abstraction）**。

------



# 我建议再优化一点（也是下一步）

当 Workspace 创建完成后，我们**不要急着接入 Connect 或 gRPC**。

我建议先做一个非常重要但容易忽略的步骤：

> **定义整个 Edge Layer 的目录规范和编码规范（Architecture Freeze V1）。**

例如：

- `cmd/` 放什么？
- `internal/` 如何分层？
- Handler、Client、DTO 分别放在哪里？
- Adapter 和 BFF 哪些目录保持一致？

这一步只需要半小时，却能避免后续目录越长越乱。**我建议我们把这套规范定下来，再开始接入 `grpc-contracts` 和实现第一个 HTTP → gRPC 接口。**我认为这样会更稳，也更符合你希望沉淀长期资产的目标。
