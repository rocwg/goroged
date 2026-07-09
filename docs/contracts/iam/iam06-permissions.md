收到，我后面会保持我们之前约定的风格：**正常段落 + 紧凑列表 + 横向表达**，不再使用那种一行只有几个字的排版。

------

## 我先评价一下这份旧的 permissions.http

我认为它的问题比前几个 Domain 更明显，主要有四点。

第一，它暴露了 UI。

例如：

```text
/menu-btn
/create-btn
/update-btn
```

这些都是站在"页面"的角度命名，而不是站在 Domain 的角度。

第二，它把 Permission 等同于 Button。

实际上，你的数据库已经说明了：

```sql
permissions
```

负责的是 **Capability（能力）**。

今天是 Button。

以后完全可能增加：

- API Permission
- Data Permission
- Operation Permission

所以 Contract 不应该写死 Button。

第三，它缺少 Resource 的基本生命周期。

目前只有：

- Create
- Update
- Delete

但是没有：

- Detail
- Query

以后权限越来越多时，就没法管理了。

第四，它没有体现 Menu 与 Permission 的关系。

你的数据库是：

```sql
permission
    └── menu_id
```

因此，Permission 是挂在 Menu 下的能力集合。

Consumer 其实需要的是：

> 查询某个菜单下有哪些 Permission。

而不是：

```text
GET /permissions/menu-btn
```

------

# 我建议冻结后的 permissions.http

```http
########################################################
# IAM · Permission（权限）
#
# Domain
# - 权限生命周期管理
# - 菜单能力管理
# - 权限编码管理
#
# Consumer Contract
# 面向 Browser / App / BFF。
# Bridge 负责映射到 PermissionService。
########################################################

@host=http://localhost:8080

########################################################
# Query（查询）
########################################################

###
# 权限详情
GET {{host}}/api/v1/iam/permissions/100

###
# 查询菜单权限
GET {{host}}/api/v1/iam/menus/5/permissions

########################################################
# Command（变更）
########################################################

###
# 创建权限
POST {{host}}/api/v1/iam/permissions
Content-Type: application/json

{
  "menuId": 5,
  "code": "sys:user:create",
  "name": "新增用户",
  "remark": ""
}

###
# 修改权限
PUT {{host}}/api/v1/iam/permissions/100
Content-Type: application/json

{
  "code": "sys:user:create",
  "name": "新增用户",
  "remark": ""
}

###
# 删除权限
DELETE {{host}}/api/v1/iam/permissions/100

########################################################
# Lookup（选择）
########################################################

###
# 权限选项查询
POST {{host}}/api/v1/iam/permissions/options/query
Content-Type: application/json

{
  "name": "",
  "menuId": 5,
  "limit": 50
}
```

------

## 为什么把菜单权限改成：

```http
GET /menus/{menuId}/permissions
```

这是我认为这一版最大的调整。

因为 Permission 本身就是 Menu 的子资源。

也就是说：

```text
Menu
 ├── Permission
 ├── Permission
 └── Permission
```

因此：

```http
GET /menus/5/permissions
```

比：

```http
GET /permissions/menu-btn
```

更符合 REST，也更符合你的 Domain。

------

## 为什么没有分页？

Permission 的数量通常不会特别大。

例如：

```text
用户管理

新增

删除

修改

导出

导入

重置密码
```

一个菜单几十个 Permission 已经很多了。

Consumer 真正需要的是：

```http
GET /menus/{id}/permissions
```

直接全部返回即可。

如果未来真的出现成百上千个 Permission，再增加分页也不迟。

------

## 为什么没有 Relationship？

因为：

Role ↔ Permission

已经放到了：

```http
PUT /roles/{id}/authorization
```

统一维护。

Permission 自己不负责授权。

它只是 Resource。

这一点与你的数据库设计完全一致。

------

## 为什么保留 Lookup？

例如：

以后可能出现：

- 创建角色模板
- 权限复制
- 权限选择器

这些页面仍然需要：

```http
POST /permissions/options/query
```

所以保留它，与其它 Domain 保持一致。

------

# 最后，我想提出一个我认为非常值得写进 Handbook 的设计原则。

经过 Identity、User、Role、Department、Menu、Permission 六个 Domain，我们已经可以总结出一个统一规律：

| Domain     | Resource | Relationship     | Authorization     | Tree |
| ---------- | -------- | ---------------- | ----------------- | ---- |
| Identity   | 身份     | —                | —                 | —    |
| User       | 用户     | Role、Department | —                 | —    |
| Role       | 角色     | User             | Menu + Permission | —    |
| Department | 部门     | User             | —                 | ✓    |
| Menu       | 菜单     | —                | —                 | ✓    |
| Permission | 权限     | —                | —                 | —    |

这六个 Contract 的边界非常清晰：**Menu 管导航、Permission 管能力、Role 管授权。** 它们各司其职，不相互越界。我认为这是你整个 IAM 设计中最有价值的地方，也是很多现有 RBAC 系统容易混淆的部分。

另外，还有一个小建议供你考虑。`GET /menus/{menuId}/permissions` 和 `POST /permissions` 并不冲突，它们分别表达了两个不同的视角：前者是**从菜单查看能力集合**，后者是**对权限资源本身进行管理**。这种"资源视角"与"父子资源视角"并存，我认为是比较自然且可扩展的设计。