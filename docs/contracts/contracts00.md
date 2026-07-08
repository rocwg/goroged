# 我觉得今天真正的新概念不是目录，而是架构语言

以前我们一直说：

> Contract。

今天，我认为应该升级成：

```ini
Contract
│
├── Consumer Contract
│      Browser/App ⇄ Edge
│
└── Service Contract
       Edge/Core ⇄ Core
```

这两个都是 Contract。

但是：

- 生命周期不同。
- 演进速度不同。
- 维护者不同。
- 面向对象不同。

所以应该明确区分。

------

# 我还想再往前走一步

其实，我们整个体系现在已经形成了一套完整的语言。

| 层次              | 第一公民                 | 代表仓库                                     |
| ----------------- | ------------------------ | -------------------------------------------- |
| Consumer Contract | HTTP / OpenAPI / `.http` | `goro-edge`                                  |
| Edge              | HTTP Adapter、BFF        | `goro-edge`                                  |
| Service Contract  | Proto                    | `grpc-contracts`                             |
| Core              | Pure gRPC Service        | `goro-dict-area-service`、`jaro-iam-service` |
| Infrastructure    | PostgreSQL、Redis、MQ    | 各 Service 内部                              |

我认为，这比单纯说"四层架构"更加准确。

因为它第一次把 **HTTP 提升为一种正式的 Contract**，但同时又没有污染 `grpc-contracts` 的纯净性。

------

## 最后，我还有一个想法（也是我认为最值得长期坚持的）

我们一直在强调：

> **Proto 是第一公民。**

今天我想把这句话稍微升级一下。

真正的第一公民，不是 **Proto 文件** 本身，而是 **Contract（契约）**。

只是这个 Contract 在不同层次有不同的表现形式：

- 面向浏览器、App、第三方调用者时，它表现为 **Consumer Contract（HTTP、OpenAPI、`.http`）**。
- 面向 Edge 与 Core、Core 与 Core 通信时，它表现为 **Service Contract（Proto）**。

因此，我建议以后你的整个架构思想可以概括为一句话：

> **Contract First Architecture（契约优先架构）。**

其中：

- `grpc-contracts` 负责维护 **Service Contract**。
- `goro-edge` 负责维护 **Consumer Contract**。

二者共同组成了整个系统的契约体系，但各自保持纯净、职责单一。我认为，这是目前我们讨论下来最自然、也最符合长期演进的一种划分。



# 我最后想到一个更重要的原则

今天，我们已经有：

```
grpc-contracts
```

它管理：Proto。

以后：**任何 Contract，都应该遵守一个原则**：

> **Contract 永远跟着它的 Provider（提供者）走，而不是跟着 Consumer（消费者）走。**

例如：

| Provider         | Contract       | 放置位置            |
| ---------------- | -------------- | ------------------- |
| `grpc-contracts` | Proto          | `grpc-contracts`    |
| `adapter`        | HTTP / OpenAPI | `adapter/contracts` |
| `bff`            | HTTP / OpenAPI | `bff/contracts`     |

这意味着：

- `grpc-contracts` 提供 gRPC 契约，所以 Proto 在那里。
- `adapter` 提供 HTTP API，所以 HTTP 契约在 `adapter`。
- `bff` 提供 HTTP API，所以 HTTP 契约在 `bff`。

这样，每一个 Contract 都和它的实现代码、生命周期、版本一起演进。

------

## 我认为，这其实和你一直坚持的理念完全一致

你一直强调：

> **我的数字资产要高内聚、低耦合。**

那么：

- **Proto** 和 `grpc-contracts` 高内聚。
- **Adapter HTTP Contract** 和 Adapter 高内聚。
- **BFF HTTP Contract** 和 BFF 高内聚。

没有任何一种 Contract 被放到一个"公共目录"等待大家共享。

**我越来越认为，这才是真正符合"Contract First Architecture"的一种组织方式**。





（1）组织方式：Contract First Architecture

（2）Browser / App ： Consumer Contract ： goro-edge ：Service Contract（Proto）: Pure gRPC Core Service:Infrastructure（DB）

（3）Pure gRPC Core Service: 两个技术栈（go\spring）

（4）2跳请求链路（浏览器 → KrakenD EE/Higress/云厂商 → 终端）：适用简单稳定API产品层；

关于 2跳请求链路，因为我的技术不足（KrakenD EE/Higress）和预算不足（云厂商），线下学习实践阶段先用 Adapter 代替。

（5）3跳请求链路（浏览器 → KrakenD CE → Go-BFF → 终端）：适用多端聚合、复杂业务逻辑、等其它需要bff的场景；

（6）1跳请求链路（浏览器 → 终端）：适用实验/临时链的场景；试水Spring、Go直连DB；本质 = innovation sandbox（创新沙盒）



计算链路分流系统（Execution Routing Architecture）有三个核心能力：

① 请求分流（Route Decision）

KrakenD 决定入口

链路选择执行路径

② 计算分层（Compute Placement）

Go：编排 / 转发 / 轻逻辑

Java：事务 / 权限 / 数据一致性

③ 生命周期管理（Lifecycle Governance）

试水 → 正式化 → 下沉/上浮



DRMA（Dual-Route Minimal Architecture）是不是不再适用