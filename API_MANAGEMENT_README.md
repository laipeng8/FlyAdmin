# API管理系统使用说明

## 概述

API管理系统用于管理系统中各个菜单对应的API接口，实现菜单与API的关联管理。系统支持API的增删改查、批量操作、重复性检查等功能。

## 功能特性

### 1. API管理
- ✅ 创建、编辑、删除API
- ✅ 批量删除API
- ✅ 支持按菜单ID查询API
- ✅ API代码和URL唯一性检查
- ✅ 与菜单关联管理

### 2. 查询功能
- ✅ 分页查询API列表
- ✅ 按代码、URL、菜单ID筛选
- ✅ 获取API详情
- ✅ 获取所有API列表

### 3. 数据验证
- ✅ API代码唯一性验证
- ✅ API地址唯一性验证
- ✅ 菜单存在性验证
- ✅ 参数格式验证

## 数据库表结构

### gc_menu_api_list (菜单API列表表)
- `id`: 主键
- `code`: API代码（唯一）
- `url`: API地址（唯一）
- `menu_id`: 关联的菜单ID
- `describe`: API描述
- `created_at`: 创建时间
- `updated_at`: 更新时间
- `deleted_at`: 删除时间（软删除）

## API接口

### 1. 获取API列表
```http
GET /api/menuApi/list?page=1&page_size=10&code=user&url=/api/user&menu_id=1
```

**查询参数：**
- `page`: 页码（默认1）
- `page_size`: 每页数量（默认10）
- `code`: API代码（模糊查询）
- `url`: API地址（模糊查询）
- `menu_id`: 菜单ID（精确查询）

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "code": "user_list",
        "url": "/api/user/list",
        "menu_id": 1,
        "describe": "获取用户列表",
        "menu": {
          "id": 1,
          "name": "用户管理",
          "path": "/user"
        },
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "size": 10
  }
}
```

### 2. 创建API
```http
POST /api/menuApi/create
Content-Type: application/json

{
  "code": "user_create",
  "url": "/api/user/create",
  "menu_id": 1,
  "describe": "创建用户"
}
```

**请求参数：**
- `code`: API代码（必填，唯一）
- `url`: API地址（必填，唯一）
- `menu_id`: 菜单ID（必填）
- `describe`: API描述（可选）

### 3. 更新API
```http
PUT /api/menuApi/update
Content-Type: application/json

{
  "id": 1,
  "code": "user_create_updated",
  "url": "/api/user/create",
  "menu_id": 1,
  "describe": "创建用户（更新后）"
}
```

**请求参数：**
- `id`: API ID（必填）
- `code`: API代码（必填，唯一）
- `url`: API地址（必填，唯一）
- `menu_id`: 菜单ID（必填）
- `describe`: API描述（可选）

### 4. 删除API
```http
DELETE /api/menuApi/delete/1
```

### 5. 批量删除API
```http
POST /api/menuApi/batch-delete
Content-Type: application/json

{
  "ids": [1, 2, 3]
}
```

### 6. 获取API详情
```http
GET /api/menuApi/get/1
```

### 7. 根据菜单ID获取API列表
```http
GET /api/menuApi/menu/1
```

### 8. 获取所有API
```http
GET /api/menuApi/all
```

### 9. 检查API代码是否存在
```http
GET /api/menuApi/check-code?code=user_list&exclude_id=1
```

**查询参数：**
- `code`: 要检查的API代码（必填）
- `exclude_id`: 排除的API ID（可选，用于更新时检查）

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "exists": false,
    "code": "user_list"
  }
}
```

### 10. 检查API地址是否存在
```http
GET /api/menuApi/check-url?url=/api/user/list&exclude_id=1
```

**查询参数：**
- `url`: 要检查的API地址（必填）
- `exclude_id`: 排除的API ID（可选，用于更新时检查）

## 使用示例

### 1. 创建用户管理相关API
```json
[
  {
    "code": "user_list",
    "url": "/api/user/list",
    "menu_id": 1,
    "describe": "获取用户列表"
  },
  {
    "code": "user_create",
    "url": "/api/user/create",
    "menu_id": 1,
    "describe": "创建用户"
  },
  {
    "code": "user_update",
    "url": "/api/user/upload",
    "menu_id": 1,
    "describe": "更新用户"
  },
  {
    "code": "user_delete",
    "url": "/api/user/delete",
    "menu_id": 1,
    "describe": "删除用户"
  }
]
```

### 2. 创建角色管理相关API
```json
[
  {
    "code": "role_list",
    "url": "/api/role/list",
    "menu_id": 2,
    "describe": "获取角色列表"
  },
  {
    "code": "role_create",
    "url": "/api/role/create",
    "menu_id": 2,
    "describe": "创建角色"
  },
  {
    "code": "role_permission",
    "url": "/api/role/permission",
    "menu_id": 2,
    "describe": "设置角色权限"
  }
]
```

### 3. 创建菜单管理相关API
```json
[
  {
    "code": "menu_list",
    "url": "/api/menu/list",
    "menu_id": 3,
    "describe": "获取菜单列表"
  },
  {
    "code": "menu_tree",
    "url": "/api/menu/tree",
    "menu_id": 3,
    "describe": "获取菜单树"
  },
  {
    "code": "menu_create",
    "url": "/api/menu/create",
    "menu_id": 3,
    "describe": "创建菜单"
  }
]
```

## 错误处理

### 常见错误码

| 错误码 | 说明 |
|--------|------|
| 400 | 参数错误 |
| 404 | API不存在 |
| 409 | API代码或地址已存在 |
| 422 | 菜单不存在 |

### 错误响应示例
```json
{
  "code": 409,
  "message": "API代码已存在",
  "data": null
}
```

## 注意事项

1. **唯一性约束**: API代码和URL在系统中必须唯一
2. **菜单关联**: 每个API必须关联到一个有效的菜单
3. **软删除**: 删除操作采用软删除，不会物理删除数据
4. **批量操作**: 批量删除时，如果某个ID不存在，会跳过该记录
5. **查询性能**: 建议合理使用分页和筛选条件，避免查询大量数据

## 权限控制

API管理系统通常需要以下权限：
- `menuApi:list` - 查看API列表
- `menuApi:create` - 创建API
- `menuApi:update` - 更新API
- `menuApi:delete` - 删除API
- `menuApi:batch-delete` - 批量删除API

## 最佳实践

1. **命名规范**: API代码使用下划线命名法，如 `user_list`
2. **URL规范**: API地址使用RESTful风格，如 `/api/user/list`
3. **描述清晰**: 为每个API提供清晰的描述信息
4. **菜单分组**: 将相关功能的API关联到同一个菜单下
5. **定期清理**: 定期清理不再使用的API记录 