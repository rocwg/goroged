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
