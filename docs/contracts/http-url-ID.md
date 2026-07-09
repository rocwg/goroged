谢谢你的信任。

我也想明确一下我的原则。

以后我们设计 **goro**，我不会以"最流行"或者"最纯 REST"作为目标，而是以：

> **十年以后，这套架构是否还值得维护。**

作为第一原则。

很多今天看起来"优雅"的设计，实际上几年以后会变成负担；而很多企业一直坚持的设计，看起来没那么"时髦"，却非常稳定。

------

## 关于 ID，我建议 goro 采用下面这套原则。

### 第一条：HTTP Resource ID 可以出现在 URL

例如：

```text
GET    /users/{id}
PUT    /users/{id}
DELETE /users/{id}

PUT    /users/{id}/roles
PUT    /users/{id}/status
```

不用因为安全而改成：

```text
POST /users/detail
```

这种 RPC 风格。

------

### 第二条：HTTP ID ≠ 数据库主键（这是重点）

这是我最希望从第一天就建立的思想。

不要认为：

```text
/users/1001
```

就是：

```sql
users.id = 1001
```

以后整个链路应该是：

```text
HTTP Resource ID
        ↓
Bridge Mapping
        ↓
Proto Identifier
        ↓
Domain Identifier
        ↓
Repository
        ↓
Database Primary Key
```

对于第一个版本来说：

可能：

```text
HTTP ID == DB ID
```

Bridge 什么也不用做。

但是以后：

数据库：

```text
BIGINT
```

改成：

```text
UUID
```

HTTP 可以一点都不用改。

这就是 Bridge 的价值。

------

### 第三条：安全永远依赖授权，不依赖隐藏 URL

例如：

管理员：

```text
PUT /users/1001/status
```

Bridge：

↓

```protobuf
ChangeUserStatus()
```

↓

Service：

↓

```java
permissionService.check(...)
```

而不是：

```text
POST /change-status
```

就觉得：

"别人不知道 URL，所以安全。"

这是企业系统的大忌。

------

## 再说一个我最近一直在思考的问题。

我觉得我们应该把 **HTTP Contract** 和 **Proto Contract** 的职责划分得更彻底。

例如：

HTTP：

```text
PUT /users/{id}/roles
```

它表达的是：

> **修改 User 与 Role 的关系。**

而 Proto：

```protobuf
rpc AssignRoles(...)
```

它表达的是：

> **执行 AssignRoles 这个业务动作。**

你会发现：

HTTP 是资源语言（Resource Language）。

Proto 是领域语言（Domain Language）。

两者**不需要一模一样**。

Bridge 正是负责完成：

```text
Resource Language → Domain Language
```

的转换。

这是我越来越确信的一件事。

------

## 关于 Query，我也想调整一下昨天的建议。

昨天我说：

```text
GET /users
```

今天结合你的实际开发经验，我觉得：

对于企业后台来说：

**Query 和 Command 应该区别对待。**

以后建议：

| 类型              | 风格                                        |
| ----------------- | ------------------------------------------- |
| Query（复杂查询） | `POST /users/query`                         |
| Resource（资源）  | `GET /users/{id}`                           |
| Create            | `POST /users`                               |
| Update            | `PUT /users/{id}`                           |
| Delete            | `DELETE /users/{id}` 或批量 `DELETE /users` |
| Relationship      | `PUT /users/{id}/roles`                     |

为什么我改变了观点？

因为企业后台几乎所有分页查询最后都会变成：

- 二十多个过滤条件
- 时间范围
- 多选数组
- 排序
- 高级搜索

如果坚持：

```text
GET /users?...
```

最后 URL 会越来越难维护。

所以：

**复杂 Query 使用 POST，我认为完全合理，而且比为了 REST 而 REST 更符合企业实践。**

------

# 最后，我想提出一个我认为会成为 goro HTTP Contract 的总原则。

我建议以后所有 Domain 都遵循下面这一套：

| 能力                             | 建议                                                         |
| -------------------------------- | ------------------------------------------------------------ |
| Query（复杂查询）                | `POST /{resources}/query`                                    |
| Get（详情）                      | `GET /{resources}/{id}`                                      |
| Create                           | `POST /{resources}`                                          |
| Update                           | `PUT /{resources}/{id}`                                      |
| Delete                           | `DELETE /{resources}/{id}`（单个）或 `DELETE /{resources}`（批量） |
| Relationship                     | `PUT /{resources}/{id}/{relationship}`                       |
| Action（极少数无法资源化的动作） | `POST /{resources}/{id}/{action}`                            |

我特意保留了最后一条 **Action**。

因为现实企业系统一定会遇到一些**无法自然资源化**的业务，例如：

- 导入 Excel
- 导出报表
- 发送验证码
- 重新计算
- 重新生成缓存
- 审批流中的"提交""驳回"

如果为了追求绝对 REST 而硬设计资源，反而会让接口变得别扭。因此，我建议 **goro** 不走 REST 教条主义，而是采用一套**资源优先（Resource First）、动作兜底（Action as Exception）**的契约规范。我认为，这既符合长期维护，也符合绝大多数企业项目的实际情况。