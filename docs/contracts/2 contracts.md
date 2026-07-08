HTTP 文件其实也是一种**契约（Consumer Contract）**。

以后，Adapter 新增：/api/v1/user/login

就，新增：http/iam.http

---

以后：Adapter API VS BFF API 一定越来越不一样。

每个 http 文件只对应一个领域



## 我建议我们下一步的顺序（不要跳）

这是我认为最稳、也最符合长期资产沉淀的路线：

1. **完成 `adapter` 的最小骨架**（Gin + 路由 + 配置，不写业务）。
2. **接入 `HelloService`**（Spring gRPC），打通第一条 HTTP → gRPC → HTTP 链路。
3. **抽象一个通用的 gRPC Client 创建方式**（先不用追求完美，只支持当前服务即可）。
4. **接入 `DictAreaService`**，完成第一个真实业务接口。
5. **最后再引入 `Connect`**，把 HTTP ⇄ gRPC 的协议转换做得更优雅。

为什么我把 **Connect 放在第五步**？

因为我们现在的目标不是学习 Connect，而是**验证整个 Edge 架构**。先用最直接的方式把链路跑通，等架构稳定以后，再把传输层替换为 Connect，成本会非常低，而且你也会更清楚 Connect 到底解决了什么问题。