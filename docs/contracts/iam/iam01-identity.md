当然，我们正式开始。

不过在开始之前，我想先**修正我昨天提出的一条规范**。

这是我经过一晚思考后的结论。

------

# 我想修正的一点

昨天我建议所有 Contract 按下面顺序排列：

```text
Query
Command
Relationship
Authorization
Lookup
Action
```

今天我觉得，这个顺序**对于大多数 Domain 是正确的**，但是 **Identity 是唯一的例外。**

原因很简单。

Identity 本身不是 CRUD Domain。

它没有：

- Create Identity
- Delete Identity

它本质上是：

> Authentication + Session + Token

所以它天然就是：

**Action First Domain。**

因此，我建议 Handbook 增加一句：

> **默认采用统一章节顺序；但对于以行为为中心（Action-Centric）的 Domain，如 Identity，可以将 Action 放在最前。**

这不是破坏规范，而是让规范更符合业务语义。

------

# Identity Domain Definition

我建议文件头先固定下来。

```http
########################################################
# IAM · Identity（身份认证）
#
# Domain
# - 身份认证（Authentication）
# - 会话管理（Session）
# - Token 生命周期管理
#
# Consumer Contract
# 面向 Browser / App / BFF。
# Bridge 负责映射到 IdentityService。
########################################################

@host=http://localhost:8080
```

这里注意。

我已经不用：

```text
Auth
```

而统一叫：

```text
Identity
```

因为：

以后：

```text
Login

Logout

Refresh

Me
```

其实：

都是：

Identity。

------

# Action（动作）

Identity 的核心。

```http
########################################################
# Action（动作）
########################################################

###
# 登录
POST {{host}}/api/v1/iam/identity/login
Content-Type: application/json

{
  "phone": "",
  "password": ""
}

###

# 登出
POST {{host}}/api/v1/iam/identity/logout

###

# 刷新访问令牌
POST {{host}}/api/v1/iam/identity/token/refresh
Content-Type: application/json

{
  "refreshToken": ""
}
```

为什么：

我仍然保留：

```text
POST /login
```

因为：

Login：

不是：

Resource。

而是：

Action。

------

# Query（查询）

这里只有一个。

```http
########################################################
# Query（查询）
########################################################

###
# 当前登录用户
GET {{host}}/api/v1/iam/identity/me
```

这里：

不要：

```text
/current-user

/profile

/info
```

统一：

```text
/me
```

行业：

基本：

已经：

统一。

------

# 为什么没有 Command？

因为：

Identity：

没有：

```text
Create

Update

Delete
```

------

# 为什么没有 Relationship？

没有。

------

# 为什么没有 Authorization？

没有。

授权：

属于：

Role。

不是：

Identity。

------

# 为什么没有 Lookup？

没有。

------

# 最终目录

Identity：

以后：

永远：

只有：

```text
Action

Query
```

是不是：

特别自然。

------

# 我还想补充两个我建议现在就冻结的小规范

## 第一条：Identity 永远使用单数

不是：

```text
/auth

/identities
```

而是：

```text
/identity
```

因为：

它不是：

资源集合。

而是：

认证上下文。

------

## 第二条：Token 不作为独立 Resource

不要：

```text
/tokens

/token/create

/token/delete
```

而是：

放在：

Identity。

例如：

```text
POST /identity/token/refresh
```

以后：

如果：

增加：

```text
Revoke

Exchange

Introspection
```

也：

全部：

放在：

```text
/identity/token/*
```

下面。

------

# 最终冻结版（V1）

我建议 `identity.http` 最终保持如下结构：

```text
IAM · Identity

Action
    POST /identity/login
    POST /identity/logout
    POST /identity/token/refresh

Query
    GET /identity/me
```

整个文件只有 **4 个接口**，没有任何冗余的 `loginByPhone`、`logoutAll`、`checkToken`、`getProfile` 等动作。

------

## 最后，我想提出一个可能会影响整个 goro Handbook 的观点。

经过我们这几天讨论，我越来越倾向于把 Domain 分成两类：

| Domain 类型                      | 代表                                     | 推荐章节顺序                                                 |
| -------------------------------- | ---------------------------------------- | ------------------------------------------------------------ |
| **Resource-Centric（资源中心）** | User、Role、Department、Menu、Permission | Query → Command → Relationship → Authorization → Lookup → Action |
| **Action-Centric（行为中心）**   | Identity                                 | Action → Query                                               |

这样，规范仍然统一，但不会为了形式而让 `identity.http` 先出现一个几乎为空的 Query 章节。**规范服务于业务，而不是业务迁就规范。**

我认为，这也是 `goro` Handbook 应该体现的一个设计理念。