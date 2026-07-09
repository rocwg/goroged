adapter：第一句话：

> Stateless HTTP ⇄ gRPC Protocol Adapter.



Architecture Rule。

第一段：

> Temporary HTTP → gRPC Protocol Bridge.

第二段：

> Replaced by Envoy / Higress / KrakenD EE in production.

这样：Bridge 这个名字：仍然成立。



明确声明：

> **Bridge 是一个协议桥（Protocol Bridge），当前用于本地开发、学习以及没有 Envoy / Higress / KrakenD Enterprise 的场景。在生产环境中，这一职责预计由专业 Gateway 或 Proxy 承担。**

这样，**名字表达职责，README 表达定位，ERA 表达架构**，三者各司其职，不会互相混淆。



#### 2026年7月9日

##### 1）Register & API 发布清单

对于 `goro-edge`，**Bridge 不应该追求"自动暴露所有 RPC"**。

Bridge 的价值，不是少写几百行代码，而是**明确、可审计地决定哪些 gRPC 能成为 HTTP API**。

所以，**Register 不仅是注册器，它实际上就是整个 Bridge 的"API 发布清单（Publish List）"**。我认为，这比任何自动生成方案都更符合你正在建设的长期工程体系。



##### 2）Bridge 当前三个建议

我建议直接作为项目规范固定下来：

1. **每个 Adapter 都保留自己的 `adapter.go`**，即使现在内容相同，也不要为了消除几十行重复代码而提前抽象。
2. **所有方法统一使用指针接收者 `func (a *Adapter)`**，不要混用 ==***值接收者***== 和 ==***指针接收者***==，这也是 Go 官方推荐的做法。
3. **构造函数统一命名为 `New`**，例如 `dictarea.New(clients)`、`hello.New(clients)`、`iam.New(clients)`，既符合 Go 的习惯，也让代码更简洁。

---

