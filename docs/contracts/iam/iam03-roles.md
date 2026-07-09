我认真看了一遍。

先给结论：

> **这一版我认为已经有 96 分了。**
>
> 这不是"旧 Spring Boot 接口"，而是一份真正意义上的 **Consumer API Contract**。

我不会大改，我只做一些**冻结前的统一**，让它以后能够成为 `goro` Handbook 的标准模板。

------

# 我建议最终结构

```text
IAM · Role

Query
Command

Relationship
    User（Increment）

Authorization
    Menu + Permission（Replace）

Lookup
```

注意这里。

我建议：

**Authorization 从 Relationship 中独立出来。**

因为它已经不是普通 Relationship。

------

# 第一处：补充文件头（与 users 保持一致）

建议统一。

```http
########################################################
# IAM · Role（角色）
#
# Domain
# - 角色生命周期管理
# - 角色基础信息管理
# - 用户与角色关系维护
# - 角色授权管理
#
# Consumer Contract
# 面向 Browser / App / BFF。
# Bridge 负责映射到 RoleService。
########################################################

@host=http://localhost:8080
```

以后：

所有 Domain：

保持一致。

------

# 第二处：Query

这里：

我建议：

只有一个小调整。

现在：

```json
{
  "page": 1,
  "size": 20,
  "name": "",
  "status": false
}
```

建议：

```json
{
  "page": 1,
  "size": 20,

  "name": "",

  "status": null
}
```

为什么？

因为：

Role：

不像：

User。

Role：

通常：

不会：

有：

keyword。

就是：

Name。

所以：

保持：

name。

很好。

另外：

status：

建议：

允许：

null。

表示：

全部。

------

# 第三处：Relationship Query

这一块。

建议：

统一：

Users。

例如：

```json
{
  "page": 1,
  "size": 20,

  "username": "",

  "phone": ""
}
```

为什么？

因为：

这里：

Consumer：

页面：

就是：

两个：

输入框。

和：

users.query：

保持：

一致。

不要：

突然：

keyword。

------

# 第四处：Relationship Command

这里：

我建议：

不要：

改。

保持：

```text
POST /users/add

POST /users/remove
```

原因：

我们：

已经：

冻结：

Relationship Increment。

以后：

Department：

完全：

复制。

------

# 第五处：Authorization

这里：

我建议：

正式：

改章节。

不要：

再：

写：

```text
Relationship

Replace
```

建议：

```
########################################################
# Authorization（授权）
#
# Replace（整体维护）
########################################################
```

因为：

Authorization：

以后：

越来越：

复杂。

菜单。

按钮。

数据权限。

API。

Scope。

都是：

Authorization。

不是：

普通：

Relationship。

------

# 第六处：Authorization Response

这里：

虽然：

http：

不用：

写：

Response。

但是：

Handbook：

建议：

冻结。

例如：

```json
{
  "menuTree": [

  ],

  "checkedMenuIds": [

  ],

  "checkedPermissionIds": [

  ]
}
```

以后：

Bridge：

↓

Proto：

都：

统一。

------

# 第七处：Lookup

这里。

昨天：

我们：

已经：

讨论：

成熟。

我建议：

正式：

补上。

```http
########################################################
# Lookup（选择）
########################################################

###
# 角色选项查询
POST {{host}}/api/v1/iam/roles/options/query
Content-Type: application/json

{
  "name": "",

  "status": true,

  "limit": 20
}
```

这里：

为什么：

不是：

```json
{
    "page":1
}
```

因为：

Lookup。

不是：

管理。

就是：

选择。

和：

users。

统一。

------

# 第八处：我建议增加一个说明（非常重要）

Role。

其实：

是：

IAM：

唯一：

同时：

拥有：

两种：

Relationship Pattern。

建议：

保留：

说明。

例如：

```
########################################################
# Relationship（关系）
#
# Increment（增量维护）
#
# 适用场景：
# Role ↔ User
########################################################
```

下面：

再：

```
########################################################
# Authorization（授权）
#
# Replace（整体维护）
#
# 适用场景：
# Role ↔ Menu
# Role ↔ Permission
########################################################
```

以后：

别人：

一看：

就：

知道：

为什么：

两个：

模式。

------

# 我建议冻结后的目录

```text
IAM · Role

Query
    POST /roles/query
    GET  /roles/{id}

Command
    POST
    PUT
    DELETE
    DELETE(batch)

Relationship
    Assigned Query
    Candidate Query
    Add
    Remove

Authorization
    GET /roles/{id}/authorization
    PUT /roles/{id}/authorization

Lookup
    POST /roles/options/query
```

------

# 最后，我想提出一个我认为应该写进 Handbook 的原则。

经过 `users.http` 和 `roles.http` 两个 Domain，我们已经可以总结出一条非常稳定的设计规律：

- **资源自身的数据**，放在 **Query / Command**。
- **资源之间的成员维护**，放在 **Relationship**。
- **资源能力的整体配置**，放在 **Authorization**。

所以：

- User ↔ Role（成员关系）→ Relationship。
- Department ↔ User（成员关系）→ Relationship。
- Role ↔ Menu / Permission（授权关系）→ Authorization。

这三者虽然底层都是数据库关联表，但在 **Consumer 的业务语义** 中完全不同，因此在 HTTP Contract 中应该明确分层。我认为这是目前整个 `goro` IAM Contract 中最值得长期坚持的一条设计原则，也会让后面的 `departments.http`、`menus.http` 和 `permissions.http` 保持一致的风格。