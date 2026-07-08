# GitHub 家族

```
rocwg/
│
├── grpc-contracts
│
├── goro-edge
│
├── goro-dict-area-service
├── goro-user-service
├── goro-menu-service
├── goro-role-service
│
├── jaro-iam-service
├── jaro-tenant-service
└── jaro-notification-service
```

你会发现规律非常强。

##### 命名体系

> 长期规范

| 层级                        | 规则                       | 示例                                      |
| --------------------------- | -------------------------- | ----------------------------------------- |
| 跨语言共享资产              | **无前缀**                 | `grpc-contracts`                          |
| Go Runtime Family           | `goro-`                    | `goro-edge`、`goro-dict-area-service`     |
| JVM Runtime Family          | `jaro-`                    | `jaro-iam-service`、`jaro-tenant-service` |
| 仓库名表达职责              | 不表达协议细节             | `*-dict-area-service`、`*-iam-service`    |
| 协议（gRPC、Connect、HTTP） | 放在实现中，不放在仓库名中 | 保持仓库名稳定                            |

我认为，这套命名最大的优点不是“好看”，而是**稳定**。未来几年，无论你把底层协议从 gRPC 扩展到 Connect，或者更换网关、升级框架，**仓库名几乎都不需要改**，这正符合你一直追求的“长期可收敛”的设计目标。



# goro-edge

我认真思考了一下，这两个想法，我的评价是不一样的。

- **（1）`goro-edge` → `goroedge`：我不推荐。**
- **（2）在 `goro-edge` 仓库内，用 `go.work` 将 HTTP Adapter 和 BFF 作为两个独立 App：我反而比较赞成。**

下面说原因



（2）HTTP Adapter 和 BFF 放一个仓库，但两个 App？

这个想法，我觉得 **比完全拆仓库更优雅**。

因为它兼顾了：

- 一个产品（Edge）
- 两个可独立部署的程序

我甚至觉得，这更符合 Go 的习惯。

------

## 我会怎么设计？

例如：

```
goro-edge/
│
├── go.work
│
├── apps/
│   ├── adapter/
│   │    ├── go.mod
│   │    └── cmd/
│   │
│   └── bff/
│        ├── go.mod
│        └── cmd/
│
├── shared/
│   ├── grpcclient/
│   ├── middleware/
│   ├── config/
│   └── dto/
│
└── README.md
```

注意：这里只有一个 GitHub 仓库。但是：实际上有两个 Go Module。

| goro-edge      |                                                              |
| -------------- | ------------------------------------------------------------ |
| Adapter        | 负责：HTTP $\rightarrow$ gRPC                                |
| BFF            | 负责：HTTP $\rightarrow$ Aggregation $\rightarrow$ Saga $\rightarrow$ ViewModel |
| 全部放 shared/ | 共享：grpc client、JWT、配置、日志                           |



------

## 为什么我喜欢这种方式？

因为：**产品边界**：还是：Edge。

但是：**部署边界**：已经：独立。

例如：

以后：docker、adapter、bff；两个镜像，完全没有问题。



------

## 甚至以后可以这样

刚开始，只有：apps/adapter，没有：bff。

等：真正需要：Dashboard。聚合。Workflow。

再：新增：apps/bff。Adapter：完全不用动。

------



## 还有一个好处

Git 提交。

例如：

今天：改：shared/grpcclient。Adapter & BFF 同时受益。

如果：两个仓库。以后：复制代码、同步 PR、同步 Tag。越来越麻烦。

------

## 我唯一建议调整的一点

你说：使用 go work；我赞成。

但是：**不要为了 go.work 而 go.work**。

真正的原因应该是：

> **Edge 是一个产品，但包含多个可独立部署的应用**。

所以，我建议：

|                      |                                  |
| -------------------- | -------------------------------- |
| Repository：产品边界 | Repository $\rightarrow$ Product |
| Workspace：开发边界  | go.work $\rightarrow$ Workspace  |
| Module：应用边界     | go.mod $\rightarrow$ Application |

也就是说：这三个边界分别解决不同的问题。

------



## 我会不会再拆成两个仓库？

目前：**不会。**

因为：HTTP Adapter。BFF。未来：一定：共享：

- JWT
- grpc client
- protobuf DTO
- middleware
- tracing
- logging
- 配置

共享比例：非常高。

一个仓库。收益：明显大于：拆仓库。

------



## 我的最终建议

我认为这是目前最均衡的方案

```
grpc-contracts                （独立仓库）
        │
        ▼
goro-edge                    （一个产品，一个仓库）
│
├── go.work
│
├── apps
│   ├── adapter              （独立 Go Module）
│   └── bff                  （独立 Go Module）
│
├── shared                   （共享基础设施）
│
└── docs
```

然后：

```
goro-dict-area-service        （独立仓库）
jaro-iam-service              （独立仓库）
```

------



## 最后，我想提出一个比目录结构更重要的建议

经过这几天的讨论，我觉得你的整个体系已经逐渐形成了三个层次：

| Contract Layer                          | Edge Layer                        | Core Layer                        |
| --------------------------------------- | --------------------------------- | --------------------------------- |
| **Contracts（契约）**：`grpc-contracts` | **Edge（边缘入口）**：`goro-edge` | **Core（领域能力）**：`*-service` |

这其实已经是一种非常清晰的架构语言。

以后我们讨论问题时，可以直接说这三个层次，而不再说"这个仓库、那个仓库"。

我认为，这会让你的整个工程体系更加统一，也更容易长期维护。

