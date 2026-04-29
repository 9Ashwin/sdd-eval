## Why

当前 Go Task Manager API 只实现了基础的 CRUD 操作（List、Create、Delete），缺少单任务查询和任务更新能力。`Task.Done` 字段已定义但无法通过 API 修改，`Task.Title` 也无法编辑。补全这些缺失的 RESTful 端点，使 API 成为一个功能完整的任务管理器。

## What Changes

- 新增 `GET /tasks/{id}` 端点，返回单个任务 JSON
- 新增 `PATCH /tasks/{id}` 端点，支持部分更新 `title` 和 `done` 字段
- Store 层新增 `GetByID(id int) (Task, bool)` 方法
- Store 层新增 `Update(id int, title *string, done *bool) (Task, bool)` 方法
- 新增对应的单元测试覆盖

## Capabilities

### New Capabilities

- `task-crud`: 完整的 Task 资源 CRUD 操作，包括单资源查询和部分字段更新

### Modified Capabilities

<!-- 无已有 spec 需要修改 -->

## Impact

- `evaluation/go/internal/store/store.go` — 新增 GetByID、Update 方法
- `evaluation/go/internal/handler/handler.go` — 新增 GET /tasks/{id}、PATCH /tasks/{id} 路由和处理函数
- `evaluation/go/internal/store/store_test.go` — 新增 GetByID、Update 测试用例
