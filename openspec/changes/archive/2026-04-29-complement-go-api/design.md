## Context

当前 Go Task Manager API 使用 `net/http` 标准库，内存存储（`sync.Mutex` 保护），已有 `GET /tasks`、`POST /tasks`、`DELETE /tasks/{id}` 三个端点。本次设计补全单资源查询和部分更新。

代码路径：`evaluation/go/internal/store/store.go`（数据层）、`evaluation/go/internal/handler/handler.go`（HTTP 层）。

## Goals / Non-Goals

**Goals:**
- 新增 `GET /tasks/{id}` 端点，返回单个 Task JSON
- 新增 `PATCH /tasks/{id}` 端点，支持部分更新 `title` 和 `done`
- Store 层新增 `GetByID` 和 `Update` 方法
- 完整的单元测试覆盖

**Non-Goals:**
- 不引入外部依赖（保持 stdlib）
- 不修改现有 Task 结构体
- 不添加认证/授权
- 不添加持久化或数据库

## Decisions

### 1. 使用 PATCH 而非 PUT

**选择 PATCH**：PATCH 用于部分更新，客户端只发送需要修改的字段。PUT 要求全量替换，对客户端不友好（需要先 GET 再修改全部字段）。

`Update` 方法使用指针参数 `*string` / `*bool`：`nil` 表示"不修改该字段"，`非 nil` 表示"修改为该值"。这是 Go 中表示 optional 字段的惯用模式，避免引入单独的 options struct。

### 2. No fields validation（至少一个字段）

PATCH 请求体如果未提供任何有效字段（或 JSON 为 `{}`），返回 400 `{"error":"no fields to update"}`。这避免了无意义的空更新调用。

### 3. Store 返回 `(T, bool)` 而非 `(*T, bool)`

`GetByID` 和 `Update` 返回值类型而非指针，与现有 `Create` 返回值风格一致。调用方拿到的是 Task 副本，避免并发修改问题（Store 已用 Mutex 保护）。

## Risks / Trade-offs

- **并发安全**：现有 Mutex 模式足够。Update 操作在 Lock 内完成查找+修改，无 TOCTOU 问题。
- **PATCH 语义限制**：本次只支持 `title` 和 `done` 字段。未来如需扩展字段，只需在 handler 的 req struct 中添加对应的指针字段即可，无需修改 Store 接口。
