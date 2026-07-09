我认为 `menus.http` 是 **整个 IAM 最特殊的一个 Domain**。

为什么？

因为它**不是权限**。

你的数据库设计已经把这件事情做对了：

```text
menus
    负责 Navigation（导航）

permissions
    负责 Capability（能力）
```

所以 **Menu 不负责授权**。

Role 才负责授权。

这意味着：

> **menus.http 不应该出现任何 Role、Permission、Authorization 的接口。**

这是它和很多 Spring Boot 后台最大的区别。

------

## 我建议冻结后的 menus.http

```http
########################################################
# IAM · Menu（菜单）
#
# Domain
# - 菜单生命周期管理
# - 导航结构管理
# - 路由信息管理
#
# Consumer Contract
# 面向 Browser / App / BFF。
# Bridge 负责映射到 MenuService。
########################################################

@host=http://localhost:8080

########################################################
# Query（查询）
########################################################

###
# 菜单详情
GET {{host}}/api/v1/iam/menus/5

###

# 菜单树
GET {{host}}/api/v1/iam/menus/tree

########################################################
# Command（变更）
########################################################

###
# 创建菜单
POST {{host}}/api/v1/iam/menus
Content-Type: application/json

{
  "parentId": 0,

  "type": "M",

  "title": "",

  "name": "",

  "path": "",

  "component": "",

  "icon": "",

  "hidden": false,

  "sortOrder": 0
}

###

# 修改菜单
PUT {{host}}/api/v1/iam/menus/5
Content-Type: application/json

{
  "parentId": 0,

  "type": "M",

  "title": "",

  "name": "",

  "path": "",

  "component": "",

  "icon": "",

  "hidden": false,

  "sortOrder": 0
}

###

# 删除菜单
DELETE {{host}}/api/v1/iam/menus/5

########################################################
# Lookup（选择）
########################################################

###
# 菜单选项查询
POST {{host}}/api/v1/iam/menus/options/query
Content-Type: application/json

{
  "title": "",

  "limit": 50
}
```

------

# 我为什么增加 Menu Detail？

建议增加：

```http
GET /menus/{id}
```

原因：

编辑流程：

```text
点击编辑

↓

查询详情

↓

修改

↓

PUT
```

和：

Department。

保持：

一致。

以后：

所有：

Resource：

都有：

```text
GET /{id}
```

------

# 为什么没有分页？

这是：

很多：

后台：

最大的：

区别。

Menu。

本身：

就是：

树。

例如：

```text
系统管理

    用户管理

    角色管理

        权限管理
```

分页：

没有：

意义。

所以：

只有：

```http
GET /menus/tree
```

------

# 为什么没有 Query？

Department：

有：

```text
POST /query
```

Menu：

没有。

因为：

Menu：

天然：

就是：

Tree。

Consumer：

真正：

需要：

Tree。

不是：

分页。

------

# 为什么增加 Lookup？

例如：

很多：

页面：

需要：

```text
请选择父菜单
```

或者：

```text
复制菜单

目标菜单
```

这种：

不是：

整个：

Tree。

可能：

就是：

搜索。

例如：

```json
{
    "title":"系统",

    "limit":20
}
```

所以：

Lookup：

建议：

保留。

------

# 我建议增加 hidden

你的 SQL：

已经：

有：

```sql
hidden boolean
```

Contract：

也：

建议：

保留。

否则：

编辑：

隐藏菜单。

还得：

补。

------

# 我建议 type 保持 M/C

例如：

```text
M

C
```

不要：

写：

```text
MENU

CATALOG
```

为什么？

Consumer：

根本：

不会：

手写。

UI：

一定：

是：

Select。

Bridge：

可以：

转换。

所以：

保持：

Domain：

Enum。

------

# 为什么没有 Authorization？

这是：

今天：

我认为：

最重要：

的一点。

很多：

RBAC：

系统：

都是：

这样：

```text
Menu

↓

Role

↓

Permission
```

导致：

Menu：

越来越：

臃肿。

而你的：

设计：

其实：

已经：

很先进。

应该：

保持。

即：

```text
Menu

负责：

Navigation
Permission

负责：

Capability
Role

负责：

Authorization
```

所以：

Role：

统一：

授权：

```http
GET /roles/{id}/authorization

PUT /roles/{id}/authorization
```

Menu：

完全：

不用：

知道：

Role。

------

# 我认为 menus.http 可以作为整个 IAM 最"纯"的 Resource

经过前面几个 Domain，我发现它们可以形成一个很清晰的分工：

| Domain     | Query       | Relationship | Authorization | Tree |
| ---------- | ----------- | ------------ | ------------- | ---- |
| Identity   | ✓           | —            | —             | —    |
| User       | ✓           | Replace      | —             | —    |
| Role       | ✓           | Increment    | ✓             | —    |
| Department | ✓           | Increment    | —             | ✓    |
| **Menu**   | ✓（Detail） | —            | —             | ✓    |
| Permission | ✓           | —            | —             | —    |

这张表也说明了一个很重要的设计思想：**不是每个 Domain 都需要所有章节，但章节顺序始终一致，按需出现。** Menu 是最纯粹的资源管理模型，它不承担成员关系，也不承担授权关系，而是专注于导航结构本身。我认为这会让整个 IAM 的 Contract 边界非常清晰，也与你当前的数据库模型高度一致。