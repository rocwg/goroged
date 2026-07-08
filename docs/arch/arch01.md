## 我的建议（也是我对昨天回答的修正）

我会把**架构职责**和**仓库职责**明确分开：

| 能力                           | 属于 Edge Execution（架构） | 建议放入 `goro-edge` 仓库 |
| ------------------------------ | --------------------------- | ------------------------- |
| Bridge（HTTP ↔ gRPC）          | ✅                           | ✅                         |
| BFF                            | ✅                           | ✅                         |
| HTTP Callback（支付、Webhook） | ✅                           | ✅（推荐）                 |
| CLI（`gorocli`）               | ✅（一种入口）               | ❌，独立仓库               |
| MQ Consumer                    | 一般属于 Core 附近          | ❌，放对应 Service         |
| Scheduler                      | 多数属于 Core               | ❌，放对应 Service         |
| Gateway（KrakenD、Envoy 等）   | ✅                           | ❌，独立部署               |

所以，**Edge Execution 是一个逻辑层，不等于一个 Git 仓库。** `goro-edge` 只是这层中的一个重要实现，专门承载 HTTP 入口相关能力（Bridge、BFF、Callback 等）。而 `gorocli`、MQ Consumer、Scheduler 等虽然也属于整个执行体系的一部分，但更适合作为独立产品或直接归属各自的 Core Service，而不是全部集中到 `goro-edge`。我认为这样职责边界会更加清晰，也更符合长期演进。

