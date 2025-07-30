# FlyAdmin 后台管理系统

## 项目概述

FlyAdmin 是一个基于 Go 语言开发的企业级后台管理系统，提供完整的文件管理、用户权限控制、定时任务调度等功能。系统采用前后端分离架构，后端基于 Gin 框架开发，具有高性能、易扩展、安全可靠的特点。

## 技术栈

### 后端技术
- **框架**: Gin (Go Web框架)
- **数据库**: MySQL + GORM (ORM框架)
- **认证**: JWT (JSON Web Token)
- **日志**: Zap (高性能日志库)
- **配置**: YAML
- **定时任务**: Cron (基于 robfig/cron)
- **事件系统**: 自定义事件分发器
- **限流**: 基于令牌桶算法
- **WebSocket**: 实时通信支持

### 前端技术
- **框架**: Vue3 + Element Plus
- **构建工具**: Vite
- **状态管理**: Pinia
- **路由**: Vue Router

## 项目结构

```
server/
├── app/                          # 应用核心代码
│   ├── bootstrap.go              # 应用启动引导
│   ├── controllers/              # 控制器层 (MVC中的C)
│   │   ├── system/               # 系统管理控制器
│   │   │   ├── userController.go     # 用户管理
│   │   │   ├── roleController.go     # 角色管理
│   │   │   ├── menuController.go     # 菜单管理
│   │   │   ├── menuApiController.go  # API管理
│   │   │   ├── departmentController.go # 部门管理
│   │   │   ├── timerController.go    # 定时任务管理
│   │   │   ├── operationLogController.go # 操作日志
│   │   │   └── commonController.go   # 通用接口
│   │   └── file/                 # 文件管理控制器
│   │       ├── fileController.go     # 文件管理
│   │       └── fileGroupController.go # 文件分组管理
│   ├── models/                   # 数据模型层 (MVC中的M)
│   │   ├── adminUser.go          # 管理员用户模型
│   │   ├── role.go               # 角色模型
│   │   ├── menu.go               # 菜单模型
│   │   ├── menuApiList.go        # API列表模型
│   │   ├── department.go         # 部门模型
│   │   ├── file.go               # 文件模型
│   │   ├── timer.go              # 定时任务模型
│   │   └── operationLog.go       # 操作日志模型
│   ├── repositorys/              # 数据访问层
│   │   ├── baseRepository.go     # 基础仓储
│   │   ├── adminRepository.go    # 用户仓储
│   │   ├── roleRepository.go     # 角色仓储
│   │   ├── systemMenuRepository.go # 菜单仓储
│   │   ├── menuApiRepository.go  # API仓储
│   │   ├── departmentRepository.go # 部门仓储
│   │   ├── fileRepository.go     # 文件仓储
│   │   ├── timerRepository.go    # 定时任务仓储
│   │   └── operationLogRepository.go # 操作日志仓储
│   ├── requests/                 # 请求结构体
│   │   ├── user.go               # 用户请求
│   │   ├── role.go               # 角色请求
│   │   ├── menu.go               # 菜单请求
│   │   ├── menuApi.go            # API请求
│   │   ├── department.go         # 部门请求
│   │   ├── file.go               # 文件请求
│   │   ├── timer.go              # 定时任务请求
│   │   └── login.go              # 登录请求
│   ├── middleware/               # 中间件
│   │   ├── jwt.go                # JWT认证
│   │   ├── cors.go               # 跨域处理
│   │   ├── permission.go         # 权限验证
│   │   ├── operationLog.go       # 操作日志
│   │   ├── limiter.go            # 限流
│   │   ├── event.go              # 事件处理
│   │   └── evnCheck.go           # 环境检查
│   ├── validators/               # 数据验证器
│   │   └── cacheCodeValidator.go # 缓存验证码验证器
│   ├── event/                    # 事件定义
│   │   ├── loginEvent.go         # 登录事件
│   │   └── testEvent.go          # 测试事件
│   └── listener/                 # 事件监听器
│       └── testListener.go       # 测试监听器
├── config/                       # 配置文件
│   ├── common.go                 # 配置结构体
│   └── config.yaml               # 配置文件
├── global/                       # 全局变量和工具
│   ├── common.go                 # 全局工具函数
│   ├── model.go                  # 全局模型
│   ├── request.go                # 全局请求
│   └── response/                 # 响应处理
│       ├── response.go           # 响应结构
│       └── err_code.go           # 错误码定义
├── initialize/                   # 初始化模块
│   ├── configInit.go             # 配置初始化
│   ├── dbInit.go                 # 数据库初始化
│   ├── zapInit.go                # 日志初始化
│   ├── eventInit.go              # 事件初始化
│   └── validatorInit.go          # 验证器初始化
├── pkg/                          # 公共包
│   ├── event/                    # 事件系统
│   │   └── eventDispatcher.go    # 事件分发器
│   ├── timer/                    # 定时任务
│   │   ├── corn.go               # 旧版定时任务
│   │   └── task_manager.go       # 新版任务管理器
│   ├── websocket/                # WebSocket
│   │   └── server.go             # WebSocket服务器
│   └── reflect/                  # 反射工具
│       └── reflect_utils.go      # 反射工具函数
├── router/                       # 路由配置
│   ├── router.go                 # 主路由
│   ├── systemApi.go              # 系统API路由
│   └── commonApi.go              # 通用API路由
├── databases/                    # 数据库脚本
│   └── gin_rbac.sql              # 数据库初始化脚本
├── upload/                       # 文件上传目录
├── logs/                         # 日志文件目录
├── main.go                       # 程序入口
├── go.mod                        # Go模块文件
├── go.sum                        # Go依赖校验文件
└── README.md                     # 项目说明
```

## 核心功能模块

### 1. 用户权限管理系统

#### 1.1 用户管理
- **功能**: 管理员用户的增删改查、登录认证、密码修改等
- **核心表**: `gc_admin_user`
- **主要接口**:
  - `POST /api/user/login` - 用户登录
  - `GET /api/user/list` - 获取用户列表
  - `POST /api/user/add` - 添加用户
  - `PUT /api/user/update` - 更新用户
  - `DELETE /api/user/del` - 删除用户
  - `GET /api/user/all` - 获取所有用户

#### 1.2 角色管理
- **功能**: 角色的增删改查、角色权限分配
- **核心表**: `gc_role`, `gc_user_role`
- **主要接口**:
  - `GET /api/role/list` - 获取角色列表
  - `POST /api/role/add` - 添加角色
  - `PUT /api/role/update` - 更新角色
  - `DELETE /api/role/del` - 删除角色
  - `POST /api/role/upMenu` - 设置角色菜单权限

#### 1.3 菜单管理
- **功能**: 动态菜单配置、菜单权限控制
- **核心表**: `gc_admin_menu`
- **主要接口**:
  - `GET /api/menu/list` - 获取菜单列表
  - `POST /api/menu/add` - 添加菜单
  - `PUT /api/menu/update` - 更新菜单
  - `DELETE /api/menu/del` - 删除菜单
  - `GET /api/menu/my` - 获取用户菜单权限

#### 1.4 API管理
- **功能**: 菜单对应的API接口管理
- **核心表**: `gc_menu_api_list`
- **主要接口**:
  - `GET /api/menuApi/list` - 获取API列表
  - `POST /api/menuApi/create` - 创建API
  - `PUT /api/menuApi/update` - 更新API
  - `DELETE /api/menuApi/delete/:id` - 删除API
  - `POST /api/menuApi/batch-delete` - 批量删除API
  - `GET /api/menuApi/get/:id` - 获取API详情
  - `GET /api/menuApi/menu/:menu_id` - 根据菜单ID获取API
  - `GET /api/menuApi/all` - 获取所有API
  - `GET /api/menuApi/check-code` - 检查API代码是否存在
  - `GET /api/menuApi/check-url` - 检查API地址是否存在

#### 1.5 部门管理
- **功能**: 组织架构管理、部门层级关系
- **核心表**: `gc_department`
- **主要接口**:
  - `GET /api/department/list` - 获取部门列表
  - `POST /api/department/add` - 添加部门
  - `PUT /api/department/update` - 更新部门
  - `DELETE /api/department/del` - 删除部门
  - `PUT /api/department/upAdmin` - 更新用户部门
  - `POST /api/department/addAdmin` - 添加用户到部门

### 2. 文件管理系统

#### 2.1 文件管理
- **功能**: 文件上传、下载、删除、分类管理
- **核心表**: `gc_files`
- **主要接口**:
  - `POST /api/file/upload` - 文件上传
  - `GET /api/file/list` - 获取文件列表
  - `PUT /api/file/edit` - 编辑文件信息
  - `DELETE /api/file/delete/:id` - 删除文件
  - `POST /api/file/batch-delete` - 批量删除文件
  - `POST /api/file/test` - 测试文件上传

#### 2.2 文件分组管理
- **功能**: 文件分类、分组层级管理
- **核心表**: `gc_file_group`
- **主要接口**:
  - `GET /api/fileGroup/list` - 获取文件分组列表
  - `POST /api/fileGroup/save` - 保存文件分组
  - `PUT /api/fileGroup/edit` - 编辑文件分组
  - `DELETE /api/fileGroup/delete/:id` - 删除文件分组
  - `GET /api/fileGroup/tree` - 获取文件分组树

### 3. 定时任务管理系统

#### 3.1 任务管理
- **功能**: 动态定时任务创建、调度、监控
- **核心表**: `gc_timer_tasks`, `gc_timer_task_logs`
- **主要接口**:
  - `POST /api/timer/start` - 启动任务管理器
  - `POST /api/timer/stop` - 停止任务管理器
  - `GET /api/timer/status` - 获取任务管理器状态
  - `GET /api/timer/task/list` - 获取任务列表
  - `POST /api/timer/task/create` - 创建定时任务
  - `PUT /api/timer/task/update` - 更新定时任务
  - `DELETE /api/timer/task/delete/:id` - 删除定时任务
  - `GET /api/timer/task/get/:id` - 获取任务详情
  - `POST /api/timer/task/execute` - 手动执行任务
  - `POST /api/timer/task/test` - 测试任务
  - `PUT /api/timer/task/toggle/:id` - 切换任务状态
  - `GET /api/timer/task/logs` - 获取任务执行日志
  - `GET /api/timer/cron/examples` - 获取Cron表达式示例

### 4. 系统管理功能

#### 4.1 操作日志
- **功能**: 系统操作记录、审计追踪
- **核心表**: `gc_operation_log`
- **主要接口**:
  - `GET /api/operationLog/list` - 获取操作日志列表

#### 4.2 通用接口
- **功能**: 文件上传、验证码、系统信息等
- **主要接口**:
  - `POST /api/common/upload` - 文件上传
  - `GET /api/common/captcha` - 获取验证码
  - `POST /api/common/captcha/verify` - 验证验证码

## 数据库设计

### 核心表结构

#### 用户权限相关表
- `gc_admin_user` - 管理员用户表
- `gc_role` - 角色表
- `gc_admin_menu` - 菜单表
- `gc_menu_api_list` - API列表表
- `gc_department` - 部门表
- `gc_user_role` - 用户角色关联表
- `gc_role_menu` - 角色菜单关联表
- `gc_user_department` - 用户部门关联表

#### 文件管理相关表
- `gc_files` - 文件表
- `gc_file_group` - 文件分组表

#### 定时任务相关表
- `gc_timer_tasks` - 定时任务表
- `gc_timer_task_logs` - 定时任务日志表

#### 系统管理相关表
- `gc_operation_log` - 操作日志表

## 配置说明

### 配置文件 (config.yaml)
```yaml
env: 'dev'                    # 环境配置
db:                          # 数据库配置
  type: 'mysql'
  host: 'localhost'
  port: '3306'
  database: 'file_manager'
  user: 'root'
  password: 'password'
  table_prefix: 'gc_'

myjwt:                       # JWT配置
  secret: 'your-secret-key'
  expires_at: 36000000

app:                         # 应用配置
  host: "http://localhost:8080"
  port: ":8080"
  uploadFile: "/upload"

rate:                        # 限流配置
  limit: 15
  burst: 15

logger:                      # 日志配置
  drive: "zap"
  path: "logs"
  size: 10
  maxAge: 3
  stdOut: true

cron:                        # 定时任务配置
  order_status_update: "0 1 * * *"
```

## 部署指南

### 1. 环境要求
- Go 1.23.1+
- MySQL 8.0+
- Redis (可选，用于缓存)

### 2. 安装步骤
```bash
# 1. 克隆项目
git clone <repository-url>
cd server

# 2. 安装依赖
go mod tidy

# 3. 配置数据库
# 修改 config.yaml 中的数据库连接信息

# 4. 初始化数据库
# 执行 databases/gin_rbac.sql

# 5. 启动服务
go run main.go
```

### 3. 生产环境部署
```bash
# 编译
go build -o FlyAdmin main.go

# 运行
./FlyAdmin
```

## API接口文档

### 认证方式
系统使用JWT进行身份认证，需要在请求头中添加：
```
Authorization: Bearer <token>
```

### 响应格式
```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

### 错误码说明
- 200: 成功
- 400: 参数错误
- 401: 未授权
- 403: 禁止访问
- 404: 资源不存在
- 409: 资源冲突
- 500: 服务器内部错误

## 权限控制

### RBAC权限模型
系统采用基于角色的访问控制(RBAC)模型：
- **用户(User)**: 系统使用者
- **角色(Role)**: 权限的集合
- **菜单(Menu)**: 功能模块
- **API**: 具体的操作接口

### 权限分配流程
1. 创建菜单和对应的API
2. 创建角色并分配菜单权限
3. 为用户分配角色
4. 用户登录后根据角色获取对应权限

## 定时任务系统

### Cron表达式格式
系统使用6位Cron表达式：`秒 分 时 日 月 星期`

### 常用示例
- `0 * * * * *` - 每分钟执行
- `0 0 * * * *` - 每小时执行
- `0 0 0 * * *` - 每天凌晨执行
- `0 0 9 * * 1-5` - 工作日9点执行

### 任务类型
- HTTP请求任务：调用指定的API接口
- 支持GET、POST、PUT、DELETE、PATCH方法
- 支持自定义请求头和请求体
- 支持重试机制和超时设置

## 文件管理特性

### 文件存储
- 支持多种文件类型：图片、视频、文档等
- 按日期和分组自动组织文件
- 支持文件分类和标签管理

### 安全特性
- 文件上传大小限制
- 文件类型验证
- 文件路径安全处理
- 访问权限控制

## 监控和日志

### 日志系统
- 使用Zap高性能日志库
- 支持结构化日志
- 日志轮转和归档
- 不同级别日志分离

### 性能监控
- 请求响应时间统计
- 数据库查询性能监控
- 内存和CPU使用监控
- 错误率统计

## 安全特性

### 认证安全
- JWT令牌认证
- 令牌过期机制
- 密码加密存储
- 登录失败限制

### 数据安全
- SQL注入防护
- XSS攻击防护
- CSRF防护
- 敏感数据加密

### 访问控制
- 基于角色的权限控制
- API级别权限验证
- 操作日志记录
- 数据访问审计

## 扩展开发

### 添加新功能模块
1. 创建数据模型 (models/)
2. 创建仓储层 (repositorys/)
3. 创建控制器 (controllers/)
4. 创建请求结构体 (requests/)
5. 添加路由配置 (router/)
6. 更新数据库迁移

### 自定义中间件
```go
func CustomMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 中间件逻辑
        c.Next()
    }
}
```

### 自定义验证器
```go
type CustomValidator struct {
    ginValidator.BaseValidator
}

func (v *CustomValidator) Validator(fl validator.FieldLevel) bool {
    // 验证逻辑
    return true
}
```

## 常见问题

### 1. 数据库连接失败
- 检查数据库配置
- 确认数据库服务运行
- 验证用户名密码

### 2. 文件上传失败
- 检查上传目录权限
- 确认文件大小限制
- 验证文件类型

### 3. 定时任务不执行
- 检查Cron表达式格式
- 确认任务状态为启用
- 查看任务执行日志

### 4. 权限验证失败
- 检查用户角色分配
- 确认菜单权限配置
- 验证API权限设置

## 更新日志

### v1.0.0 (2024-01-01)
- 初始版本发布
- 基础用户权限管理
- 文件管理功能
- 定时任务系统

### v1.1.0 (2024-01-15)
- 新增API管理功能
- 优化定时任务系统
- 增强文件分组管理
- 完善操作日志

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交代码
4. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证，详见 [LICENSE](LICENSE) 文件。

## 联系方式

- 项目地址: [GitHub Repository]
- 问题反馈: [Issues]
- 邮箱: [2195143506@qq.com]

---

**注意**: 本文档会随着项目的发展持续更新，请关注最新版本。 

