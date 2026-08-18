我觉得**需要，而且这是你这个架构最容易被忽略的一部分。**

因为你目前定义的三个链路：

- 1 跳：Browser → Core
- 2 跳：Browser → Gateway → Core
- 3 跳：Browser → Gateway → BFF → Core

它们都有一个共同点：

> **都是 Browser 发起的同步请求**。

但是现实项目中，浏览器流量往往只占一部分。

------

## 我建议把入口（Entry）分类，而不是只按跳数分类

例如：

```
                    Entry
                      │
      ┌───────────────┼────────────────┐
      │               │                │
      ▼               ▼                ▼
 Browser/API      System/API      Async/Event
```

这样整个体系会完整得多。



# goro-edge

Edge Layer for Pure gRPC Services.

Contains:

- HTTP Adapter
- Backend For Frontend (BFF)

```powershell
go work sync
go test ...
go run
```

