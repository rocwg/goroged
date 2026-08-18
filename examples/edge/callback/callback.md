# 第二类：System Callback（我认为必须加入）

例如：

```
微信支付
支付宝
Stripe
物流
短信
邮件
ERP
WMS
第三方开放平台
```

它们不会经过浏览器。

例如：

1. Alipay → Callback Adapter → Core Service
2. WeChat Pay → Gateway → Callback Adapter → Core Service

这里的 Adapter 与 BFF 不同。



职责通常是：

- 验签
- IP 白名单
- 幂等
- DTO 转换
- 重试
- 限流

然后调用 Core。

例如支付成功：

```
支付平台
↓
callback
↓
verify signature
↓
convert dto
↓
grpc
↓
order service
```

这种链路几乎所有企业都会有。
