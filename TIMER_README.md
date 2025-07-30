# 定时任务系统使用说明

## 概述

本项目实现了一个完整的定时任务管理系统，可以动态设置定时启动任何已写好的接口方法。系统支持多种HTTP请求方法，具有完整的任务管理、执行日志、重试机制等功能。

## 功能特性

### 1. 任务管理
- ✅ 创建、编辑、删除定时任务
- ✅ 启用/禁用任务
- ✅ 支持多种HTTP方法（GET、POST、PUT、DELETE、PATCH）
- ✅ 自定义请求头和请求体
- ✅ 设置超时时间和重试机制

### 2. 调度功能
- ✅ 支持标准Cron表达式（6位，包含秒）
- ✅ 动态添加/移除任务
- ✅ 任务状态实时监控
- ✅ 自动重试失败的任务

### 3. 日志记录
- ✅ 详细的任务执行日志
- ✅ 执行时间统计
- ✅ 成功/失败次数统计
- ✅ HTTP响应状态码记录

### 4. 系统管理
- ✅ 任务管理器启动/停止
- ✅ 任务测试功能
- ✅ Cron表达式示例
- ✅ 分页查询和搜索

## API接口

### 任务管理器控制

#### 启动任务管理器
```http
POST /api/timer/start
```

#### 停止任务管理器
```http
POST /api/timer/stop
```

#### 获取任务管理器状态
```http
GET /api/timer/status
```

### 任务管理

#### 创建定时任务
```http
POST /api/timer/task/create
Content-Type: application/json

{
  "name": "测试任务",
  "description": "这是一个测试任务",
  "cron_expression": "0 */5 * * * *",
  "target_url": "http://localhost:8080/api/user/list",
  "method": "GET",
  "headers": "{\"Authorization\": \"Bearer token\"}",
  "body": "",
  "timeout": 30,
  "retry_count": 3,
  "retry_interval": 60
}
```

#### 更新定时任务
```http
PUT /api/timer/task/update
Content-Type: application/json

{
  "id": 1,
  "name": "更新后的任务",
  "description": "更新后的描述",
  "cron_expression": "0 0 2 * * *",
  "target_url": "http://localhost:8080/api/system/health",
  "method": "POST",
  "headers": "{\"Content-Type\": \"application/json\"}",
  "body": "{\"key\": \"value\"}",
  "status": 1,
  "timeout": 60,
  "retry_count": 5,
  "retry_interval": 120
}
```

#### 删除定时任务
```http
DELETE /api/timer/task/delete/1
```

#### 获取任务详情
```http
GET /api/timer/task/get/1
```

#### 获取任务列表
```http
GET /api/timer/task/list?page=1&page_size=10&name=测试&status=1
```

#### 手动执行任务
```http
POST /api/timer/task/execute
Content-Type: application/json

{
  "task_id": 1
}
```

#### 测试任务
```http
POST /api/timer/task/test
Content-Type: application/json

{
  "target_url": "http://localhost:8080/api/user/list",
  "method": "GET",
  "headers": "{\"Authorization\": \"Bearer token\"}",
  "body": "",
  "timeout": 30
}
```

#### 切换任务状态
```http
PUT /api/timer/task/toggle/1
```

### 日志管理

#### 获取任务执行日志
```http
GET /api/timer/task/logs?task_id=1&page=1&page_size=20
```

### 工具接口

#### 获取Cron表达式示例
```http
GET /api/timer/cron/examples
```

## Cron表达式格式

系统使用6位Cron表达式，格式为：`秒 分 时 日 月 星期`

### 常用示例

| 描述 | 表达式 | 说明 |
|------|--------|------|
| 每分钟执行 | `0 * * * * *` | 每分钟的第0秒执行 |
| 每小时执行 | `0 0 * * * *` | 每小时的第0分0秒执行 |
| 每天凌晨执行 | `0 0 0 * * *` | 每天凌晨0点0分0秒执行 |
| 每周一执行 | `0 0 0 * * 1` | 每周一凌晨0点0分0秒执行 |
| 每月1号执行 | `0 0 0 1 * *` | 每月1号凌晨0点0分0秒执行 |
| 每5分钟执行 | `0 */5 * * * *` | 每5分钟执行一次 |
| 每天上午9点和下午6点执行 | `0 0 9,18 * * *` | 每天上午9点和下午6点执行 |
| 工作日执行 | `0 0 9 * * 1-5` | 周一到周五上午9点执行 |

## 数据库表结构

### gc_timer_tasks (定时任务表)
- `id`: 主键
- `name`: 任务名称
- `description`: 任务描述
- `cron_expression`: Cron表达式
- `target_url`: 目标URL
- `method`: HTTP方法
- `headers`: 请求头（JSON格式）
- `body`: 请求体
- `status`: 状态（1:启用 0:禁用）
- `last_run_time`: 最后执行时间
- `next_run_time`: 下次执行时间
- `run_count`: 执行次数
- `success_count`: 成功次数
- `fail_count`: 失败次数
- `timeout`: 超时时间（秒）
- `retry_count`: 重试次数
- `retry_interval`: 重试间隔（秒）
- `creator`: 创建者ID

### gc_timer_task_logs (任务执行日志表)
- `id`: 主键
- `task_id`: 任务ID
- `status`: 执行状态（1:成功 0:失败）
- `message`: 执行结果消息
- `duration`: 执行时长（毫秒）
- `response`: 响应内容
- `status_code`: HTTP状态码
- `run_time`: 执行时间

## 使用示例

### 1. 创建定时清理日志任务
```json
{
  "name": "清理系统日志",
  "description": "每天凌晨2点清理7天前的系统日志",
  "cron_expression": "0 0 2 * * *",
  "target_url": "http://localhost:8080/api/system/clean-logs",
  "method": "POST",
  "headers": "{\"Content-Type\": \"application/json\"}",
  "body": "{\"days\": 7}",
  "timeout": 300,
  "retry_count": 3,
  "retry_interval": 300
}
```

### 2. 创建数据备份任务
```json
{
  "name": "数据库备份",
  "description": "每周日凌晨3点进行数据库备份",
  "cron_expression": "0 0 3 * * 0",
  "target_url": "http://localhost:8080/api/system/backup",
  "method": "POST",
  "headers": "{\"Content-Type\": \"application/json\"}",
  "body": "{\"type\": \"full\"}",
  "timeout": 1800,
  "retry_count": 2,
  "retry_interval": 600
}
```

### 3. 创建健康检查任务
```json
{
  "name": "系统健康检查",
  "description": "每5分钟检查系统健康状态",
  "cron_expression": "0 */5 * * * *",
  "target_url": "http://localhost:8080/api/system/health",
  "method": "GET",
  "headers": "",
  "body": "",
  "timeout": 30,
  "retry_count": 3,
  "retry_interval": 60
}
```

## 部署说明

### 1. 数据库初始化
执行 `databases/timer_tables.sql` 文件创建相关表结构。

### 2. 应用启动
系统会在启动时自动启动定时任务管理器，加载所有启用的任务。

### 3. 配置说明
- 任务管理器会在应用启动时自动启动
- 所有启用的任务会自动加载到调度器中
- 任务执行日志会保存到数据库中
- 支持动态添加、修改、删除任务

## 注意事项

1. **Cron表达式**: 使用6位格式，包含秒字段
2. **URL格式**: 目标URL必须是完整的HTTP/HTTPS地址
3. **超时设置**: 根据任务复杂度合理设置超时时间
4. **重试机制**: 失败任务会自动重试，注意设置合理的重试间隔
5. **日志清理**: 建议定期清理过期的执行日志
6. **权限控制**: 确保目标接口有适当的访问权限

## 故障排除

### 常见问题

1. **任务不执行**
   - 检查任务状态是否为启用
   - 验证Cron表达式格式
   - 检查目标URL是否可访问

2. **任务执行失败**
   - 查看执行日志了解具体错误
   - 检查网络连接
   - 验证请求参数格式

3. **任务管理器启动失败**
   - 检查数据库连接
   - 验证表结构是否正确
   - 查看应用日志

### 日志查看
- 应用日志: `logs/` 目录
- 任务执行日志: 通过API接口查询
- 数据库日志: 查看数据库错误日志 