

# 重新规划整个 contracts/http

```
bridge/
└── contracts/
    └── http/
        └── iam/
            ├── auth.http
            ├── users.http
            ├── roles.http
            ├── departments.http
            ├── menus.http
            ├── permissions.http
            └── README.md
```

1. README.md：不是介绍接口，而是 IAM Requirement
2. users.http：就是 User Requirement（Executable）

这样以后：README → users.http → user.proto → User Bridge → UserService

整条链路就是完全一致的。

------

## （2）下一个是不是 user？

**我的答案是：是，而且必须是 User。**

我建议整个顺序固定下来，这个顺序不是随便排的，而是有依赖关系。

| 1）Identity   | identity.http    | 例如：                                       |
| ------------- | ---------------- | -------------------------------------------- |
| 2）User       | users.http       | （1）登录（Identity）最终得到的是一个 User。 |
| 3）Role       | roles.http       | （2）User 可以拥有多个 Role。                |
| 4）Department | departments.http | （3）User 可以属于 Department。              |
| 5）Menu       | menus.http       | （4）Role 决定 Menu。                        |
| 6）Permission | permissions.http | （5）Menu 再关联 Permission。                |

所以从上到下，业务关系是连贯的。

------

## 不过，我建议在开始 `users.http` 之前，先做一件事情。

**不要急着写接口，先定义 User Domain 的职责。**

这是我现在最想帮助你做的，因为它会影响后面的 Proto。

例如，我不会再沿用旧系统那个 User。

我会重新定义：

```
User（用户）

职责：

✔ 用户生命周期管理
✔ 用户基础资料
✔ 用户状态
✔ 用户所属部门
✔ 用户角色关系

不负责：

✘ 登录（Identity）
✘ 菜单授权（Role）
✘ 权限计算（Permission）
✘ JWT Token
✘ Refresh Token
```

这一步非常重要。

以后我们写：

```
GET /users

POST /users

PUT /users/{id}

PUT /users/{id}/roles
```

都会非常自然。

------

## 我还有一个建议，我认为它会让整个 goro 的目录更加统一。

我们前面已经决定：

```
identity.http
users.http
roles.http
departments.http
menus.http
permissions.http
```

我建议 **Proto、Bridge、Service 全部采用相同命名**。

也就是说，**Contract 名称、Proto 名称、Bridge 名称、Service 名称始终保持一一对应**，这样开发时几乎不用思考“这个文件到底对应哪个 Proto、哪个 Bridge”。整个链路始终保持一致：

```
Requirement → identity.http → identity.proto → IdentityBridge → IdentityService
Requirement → users.http    → user.proto     → UserBridge     → UserService
Requirement → roles.http    → role.proto     → RoleBridge     → RoleService
```

这里我有一个**唯一的小调整建议**。

虽然 REST 风格喜欢资源复数（`/users`、`/roles`），但 **Proto 和 Service 我更建议保持单数**，例如：

```
users.http          （HTTP 资源：/users）
user.proto          （一个 User 领域）
UserService
UserBridge
```

同样：

```
roles.http   → role.proto   → RoleService
menus.http   → menu.proto   → MenuService
```

原因是：**HTTP 表示资源集合，所以用复数；Proto 和 Service 表示领域能力，所以用单数。**

这是很多成熟框架（包括 Kubernetes API、Google API）都比较常见的一种长期稳定的命名方式。我个人认为，这种区分会比全都统一成复数更加自然，也更容易阅读。

------

## 最后，我想提出一个我昨天没有想到、但今天看到你的 `.http` 后冒出来的想法。

你最初提出的是：

> **需求驱动契约，契约驱动实现，实现验证架构，架构沉淀规范。**

现在我觉得可以把第一句再精确一点，变成：

> **需求驱动可执行契约（Executable Requirement），可执行契约驱动 Provider Contract，Provider Contract 驱动实现，实现验证架构，架构沉淀规范。**

这里最大的变化是：**`.http` 不再只是一个调试文件，而是需求本身的载体**。它既是开发入口，也是验收入口；既服务于消费者（Consumer），又成为 Bridge 和 Proto 设计的依据。我认为，这会是 **goro-edge** 与很多 Contract First 实践最大的不同——你不是从 Proto 开始，而是从一份可以直接运行、可以持续验证的需求开始。这个思路，我认为已经有资格成为 **《goro-edge Architecture Handbook》** 的核心理念之一。



# 最后，我想提一个我认为值得现在就冻结的命名规范。

我建议以后所有 `.http` 文件都遵循统一章节顺序：

```
Query（查询）
Command（变更）
Relationship（关系）
Authorization（授权，可选）
Lookup（选择，可选）
Action（动作，可选）
```

不是每个 Domain 都需要六个章节，但**顺序始终一致**。例如：

- `users.http`：Query → Command → Relationship → Lookup。
- `roles.http`：Query → Command → Relationship → Authorization → Lookup。
- `menus.http`：Query → Command。
- `identity.http`：Action（login、logout、refresh）→ Query（me）。

这样，开发者打开任何一个 Contract，都能立即知道各类接口应该放在哪里，整个 `goro` 的 Handbook 和实际代码会保持高度一致。我认为，这比单纯统一 URL 风格更重要，因为它统一的是**整个团队阅读和设计 Contract 的思维方式**。

因为经过这几轮讨论，我认为你的页面需求是真实存在的，这已经不是"数据库字段泄漏"，而是**页面交互模型（Interaction Model）**。

## 我最后再补充一句我认为可以写进 Handbook 的原则

我建议把我们这几天讨论沉淀成下面这四条：

1. **Requirement 决定 Consumer API Contract，而不是数据库。**
2. **Consumer API Contract 描述页面真实的输入、选择和操作。**
3. **Bridge 负责将 Consumer Language 翻译为 Provider Language。**
4. **Provider（Proto、Domain、SQL）可以自由演进，只要 Consumer API Contract 不变。**

我认为，这四条比"用 `keyword` 还是 `username`"更重要，因为它们定义的是整个 **Contract-First** 方法论的边界。实际上，`username`、`phone`、`departmentIds` 这些字段并没有违反抽象原则，它们本身就是 **Consumer 的语言**，真正需要隐藏的是 `LIKE`、`JOIN`、索引、SQL 等 **Provider 的语言**。这两者一定要区分开。
