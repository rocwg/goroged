我认为 `departments.http` 是 **整个 IAM 最容易设计歪的一个 Domain**。

很多系统把 Department 当成普通 CRUD。

实际上，它至少承担三种职责：

1. **组织机构管理（Department Lifecycle）**
2. **组织树查询（Organization Tree）**
3. **部门成员维护（Department ↔ User）**

所以，它的 Contract 不应该只是把 `roles.http` 换个名字。

下面是我建议冻结的版本。

------

```http
########################################################
# IAM · Department（部门）
#
# Domain
# - 部门生命周期管理
# - 部门基础信息管理
# - 组织架构管理
# - 部门成员关系维护
#
# Consumer Contract
# 面向 Browser / App / BFF。
# Bridge 负责映射到 DepartmentService。
########################################################

@host=http://localhost:8080

########################################################
# Query（查询）
########################################################

###
# 部门分页查询
POST {{host}}/api/v1/iam/departments/query
Content-Type: application/json

{
  "page": 1,
  "size": 20,

  "name": "",

  "status": null
}

###

# 部门详情
GET {{host}}/api/v1/iam/departments/5

###

# 部门树
GET {{host}}/api/v1/iam/departments/tree

###

# 根据编码查询部门
GET {{host}}/api/v1/iam/departments/by-code?code=zz-01

########################################################
# Command（变更）
########################################################

###
# 创建部门
POST {{host}}/api/v1/iam/departments
Content-Type: application/json

{
  "parentId": 1,

  "name": "",

  "remark": ""
}

###

# 修改部门
PUT {{host}}/api/v1/iam/departments/5
Content-Type: application/json

{
  "parentId": 1,

  "name": "",

  "remark": ""
}

###

# 删除部门（单个）
DELETE {{host}}/api/v1/iam/departments/5

###

# 批量删除部门
DELETE {{host}}/api/v1/iam/departments
Content-Type: application/json

{
  "departmentIds": [
    5,
    6
  ]
}

########################################################
# Relationship（关系）
#
# Increment（增量维护）
#
# Department ↔ User
########################################################

#################
# Relationship Query
#################

###
# 查询部门成员
POST {{host}}/api/v1/iam/departments/5/users/query
Content-Type: application/json

{
  "page": 1,
  "size": 20,

  "username": "",

  "phone": ""
}

###

# 查询可加入部门的用户
POST {{host}}/api/v1/iam/departments/5/candidate-users/query
Content-Type: application/json

{
  "page": 1,
  "size": 20,

  "username": "",

  "phone": ""
}

#################
# Relationship Command
#################

###
# 添加部门成员
POST {{host}}/api/v1/iam/departments/5/users/add
Content-Type: application/json

{
  "userIds": [
    3,
    4
  ]
}

###

# 移除部门成员
POST {{host}}/api/v1/iam/departments/5/users/remove
Content-Type: application/json

{
  "userIds": [
    3,
    4
  ]
}

########################################################
# Lookup（选择）
########################################################

###
# 部门选项查询
POST {{host}}/api/v1/iam/departments/options/query
Content-Type: application/json

{
  "name": "",

  "status": true,

  "limit": 20
}
```

------

## 我为什么增加了两个接口？

### ① `GET /departments/{id}`

以前没有。

我建议保留。

原因很简单：

编辑页面：

```
点击编辑

↓

查询详情

↓

回显

↓

修改

↓

PUT
```

这是最标准的 Consumer 流程。

------

### ② `GET /departments/tree`

这是我认为必须增加的。

因为 Department 本身就是树。

以后很多页面都会使用：

```
组织机构

研发中心
    Java组
    Go组

行政中心

财务中心
```

如果没有 Tree。

以后：

所有页面：

都会：

自己：

拼。

------

## Lookup 为什么仍然存在？

很多地方：

不是：

组织树。

例如：

```
负责人：

请选择部门
```

或者：

```
调岗

目标部门：
```

可能：

就是：

搜索。

所以：

仍然建议：

保留：

```
POST /departments/options/query
```

和：

`users`、

```
roles
```

统一。

------

# 最后，我想提出一个我认为值得冻结的地方。

## Department 的 Query 有两类

不是一类。

而是：

```
Department Resource

↓

分页查询
详情

↓

Department Hierarchy

↓

Tree
```

所以：

以后 Handbook 我建议明确区分：

```
Resource Query
```

和：

```
Hierarchy Query
```

不是所有 Domain 都有 Tree。

但是：

只要是：

```
Department

Menu

Area

Organization
```

这种具有层级结构的 Resource，我建议统一提供一个稳定的 `GET /{resource}/tree` 接口。这不是为了前端方便，而是因为**树本身就是该领域模型的一部分**，它属于 Consumer 真正需要的业务视图，而不是 Provider 的内部实现。

------

我还想给你一个建议，供后面 `menus.http` 设计时参考：**Department 和 Menu 会成为整个 IAM 中风格最接近的两个 Domain。** 它们都是树结构资源，都有 `tree` 查询，也都有普通 CRUD，但 Department 的 Relationship 是 **Department ↔ User**，而 Menu 则没有成员关系，而是更多参与 Role 的 Authorization。这两份 Contract 保持统一风格，整个 IAM 的一致性会非常高。