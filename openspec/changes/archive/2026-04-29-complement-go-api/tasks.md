## 1. Store Layer

- [x] 1.1 实现 `Store.GetByID(id int) (Task, bool)` 方法
- [x] 1.2 实现 `Store.Update(id int, title *string, done *bool) (Task, bool)` 方法

## 2. Handler Layer

- [x] 2.1 注册 `GET /tasks/{id}` 路由并实现 getByID handler
- [x] 2.2 注册 `PATCH /tasks/{id}` 路由并实现 update handler

## 3. Tests

- [x] 3.1 添加 `TestGetByID_Exists` 和 `TestGetByID_NotFound` 测试
- [x] 3.2 添加 `TestUpdate_Title`、`TestUpdate_Done`、`TestUpdate_Both`、`TestUpdate_NotFound`、`TestUpdate_NoFields` 测试
- [x] 3.3 运行 `go test ./...` 验证全部测试通过
