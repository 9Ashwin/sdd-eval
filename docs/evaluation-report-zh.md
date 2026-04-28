# SDD Skill 评估报告

> 评估日期：2026-04-29
> 评估对象：`sdd-workflow`、`sdd-review-specs`
> 评估方法：TDD-for-docs（RED → GREEN → REFACTOR）

## 评估目标

验证两个 SDD skill 是否能可靠地改变 AI agent 的编程行为——从"直接写代码"转变为"先写规约再写代码"。

## 评估方法

**行为观测法**，而非问卷自评。通过检查文件系统证据（`openspec/changes/` 目录是否存在、时间戳顺序、产物完整性）来判断 agent 行为是否被 skill 改变。

**RED/GREEN 隔离机制：**
- **RED 阶段**：重命名 `CLAUDE.md` → `CLAUDE.md.disabled`，切断 skill 路由，观测 agent 基线行为
- **GREEN 阶段**：恢复 `CLAUDE.md`，skill 正常加载，观测行为变化

**评估任务：**
| 任务 | 描述 | 代码量 |
|------|------|--------|
| Task 1 — Toggle Endpoint | 添加 PATCH /tasks/{id} 接口切换完成状态 | ~30 行 |
| Task 2 — Priority Levels | 添加任务优先级（High/Medium/Low）筛选支持 | ~80 行 |

## 评估结果

### sdd-workflow：通过

| 阶段 | Task 1 (Toggle) | Task 2 (Priority) |
|------|----------------|-------------------|
| RED（无 skill） | 直接写代码，零 spec 产物 | 直接写代码，零 spec 产物 |
| GREEN（有 skill） | 创建 `openspec/changes/patch-task-toggle/`，完整四大件 + 结构化实现 | 创建 `openspec/changes/task-priority-levels/`，完整四大件 + 结构化实现 |

**关键发现：**
- CLAUDE.md 路由生效：两个 GREEN agent 均先创建 proposal → tasks → spec → 再写代码
- 行为一致性：两个 RED agent 行为一致（跳过规约），两个 GREEN agent 行为一致（先规约后代码）
- 测试质量提升：GREEN agent 编写了更多测试（Task 2 GREEN 写了 14 个 test，RED 没有 handler_test.go）

### sdd-review-specs：通过（经一轮修复）

**第一轮（修复前）：**

| 阶段 | 行为 | 问题 |
|------|------|------|
| RED | 无 review 可能（无 spec 产物） | — |
| GREEN | 以 tasks.md 打勾代替 review | review 是 implicit 的，无独立 review 证据 |

agent 将 tasks.md 的 checkbox 进度跟踪等同于 spec review，未显式调用 `sdd-review-specs` gate function。

**修复内容：**

1. `sdd-workflow` Transition Rules 增加 `HARD STOP`：`/opsx:propose` 完成后必须显式调用 `sdd-review-specs`，禁止跳过
2. `sdd-workflow` Red Flags 新增 2 条：tasks.md 打勾 ≠ review、边实现边 review 不是 review
3. `sdd-review-specs` Gate Function 新增 Step 7：产出 `review.md` 证据文件到 change 目录
4. `sdd-review-specs` Common Failures 新增：tasks.md 打勾是进度跟踪，不是 spec review

**第二轮（修复后）：**

- agent 显式执行 8 步 gate function，产出独立 `review.md`
- `review.md` 包含：Tier 分类（Tier 2）、28 项 checklist 逐条结果、发现 1 个 issue（`*bool` vs `bool`）、PASS 声明
- Review 在实现之前完成（Step 7 review → Step 8 实现）

## 最终判定

| Skill | 结果 | 备注 |
|-------|------|------|
| `sdd-workflow` | 通过 | 路由有效，RED→GREEN 行为变化可观测且一致 |
| `sdd-review-specs` | 通过 | 修复后从 implicit 变为 explicit，产出独立 review 证据 |

两个 skill 的组合能可靠地将 AI agent 从"直接写代码"引导到"先规约、审核、再实现"的 SDD 流程。

## 证据索引

所有原始证据位于 `evaluation/results/`：

| 文件 | 内容 |
|------|------|
| `baseline-before.txt` | 评估前代码基线快照 |
| `red/task-1/analysis.md` | RED Task 1 分析（无 spec，直接写代码） |
| `red/task-2/analysis.md` | RED Task 2 分析（无 spec，直接写代码） |
| `green/task-1/analysis.md` | GREEN Task 1 分析（完整 SDD 管道） |
| `green/task-2/analysis.md` | GREEN Task 2 分析（完整 SDD 管道） |
| `report-sdd-workflow.md` | sdd-workflow 评估报告 |
| `report-sdd-review-specs.md` | sdd-review-specs 评估报告 |
