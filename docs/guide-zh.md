# SDD 工作流指南

> 结合 OpenSpec（规约层）与 Superpowers（执行纪律层）的规约驱动开发（SDD）工作流。通用版，不绑定语言或框架。

## 核心理念

**写第一行代码之前，确保人和 AI 对"要做什么"达成了精确共识。** 规约以文件形式存在于项目中，而非散落在聊天记录中。OpenSpec 管理"是什么"，Superpowers 约束"怎么做"。

## 速查：一条变更的全生命周期

```
需求进来 → sdd-workflow 分类路由

  单行修复？           → 直接改，go test 验证
  Bug 原因不明？        → @systematic-debugging
  新功能 / 行为变更？   → 进入 SDD 管道 ↓

  /opsx:propose <名称>   → 生成四大件
  sdd-review-specs       → 审核四大件（Tier 0/1/2）
  @writing-plans         → 拆分为 2-5 分钟子任务
  /opsx:apply + @tdd     → 流水线执行
  @verification          → 验证证据
  /opsx:archive <名称>   → delta 合并 + 归档
```

## 阶段一：编写规约

### 四大件

`/opsx:propose <名称>` 在 `openspec/changes/<名称>/` 下生成四个文件：

| 产物 | 回答的问题 | 审核重点 |
|------|-----------|---------|
| `proposal.md` | 为什么做 + 范围边界 | in/out scope 是否明确？有无"and related features"这种 scope creep？ |
| `specs/` | 行为契约 | 正常路径和错误路径是否都覆盖了？是否有模糊的"should handle errors"？ |
| `design.md` | 技术方案 | 为什么选这个方案？拒绝了哪些替代方案？是否引入新依赖？ |
| `tasks.md` | 实现清单 | 每个任务是否可独立执行？是否配了测试任务？是否有"add error handling"占位符？ |

按 **为什么 → 是什么 → 怎么做 → 哪些步骤** 的顺序阅读和审核。

### 审核（sdd-review-specs）

AI 生成的规约是初稿，不是最终合同。必须人工审核后才能开始编码。

| Tier | 适用 | 范围 | 时间 |
|------|------|------|------|
| Tier 0 — 跳过 | 拼写修正、日志行、注释 | 无规约，直接 build/lint | 0 min |
| Tier 1 — 轻量 | 单字段添加、配置调整 | proposal 范围 + tasks 可执行性，扫一眼 design | 5-10 min |
| Tier 2 — 完整 | 新功能、跨包重构、架构变更 | 四大件全部 checklist | 15-30 min |

**不确定时默认 Tier 2。** 过度审核比漏掉问题便宜。

## 阶段二：执行

### 粒度细化（@writing-plans）

将 tasks.md 的粗粒度任务拆成 2-5 分钟一个的子任务。每个子任务包含：具体代码、测试命令、预期输出、commit message。

### TDD 流水线（/opsx:apply + @test-driven-development）

`/opsx:apply` 是调度器，`@test-driven-development` 是执行器：

```
对每个子任务循环：
  RED    → 写一个必定失败的测试
  GREEN  → 写最少代码让测试通过
  REFACTOR → 逐步添加真实逻辑，保持绿灯
  全部通过 → 标记完成 → 下一个
```

### 遇到失败？

不要试错。`@systematic-debugging` — 根因调查优先。

## 阶段三：交付

### 验证（@verification-before-completion）

**铁律：没有新鲜出炉的验证证据，不得声称完成。** 重新跑测试——不信任缓存结果。

### 归档（/opsx:archive）

三件事：变更目录迁移到 archive/、delta specs 合并到永久规约库、更新 project.md。

## 工具选择矩阵

| 场景 | 用这个 | 不用那个 | 原因 |
|------|--------|---------|------|
| 读已有代码 | `/opsx:explore` | `@brainstorming` | Explore 读代码；brainstorming 生成想法 |
| 从零设计新功能 | `@brainstorming` | `/opsx:explore` | Brainstorming 比较方案；explore 描述现状 |
| 生成规约 | `/opsx:propose` | `@writing-plans` | Propose 生成四大件结构；writing-plans 细化粒度 |
| 细化任务粒度 | `@writing-plans` | 手动 | Writing-plans 将粗任务转化为 2-5min 单元 |
| 执行任务 | `/opsx:apply` + `@tdd` | 单独使用 | Apply 调度；TDD 执行。流水线组合 |
| 调试失败 | `@systematic-debugging` | 直接修 | 根因调查优先，禁止试错 |
| 代码审查 | `@requesting-code-review` | "看着没问题" | 结构化审查，独立上下文 |
| 声明完成 | `@verification-before-completion` | "应该能用了" | 需要新鲜验证证据 |
| 归档 | `/opsx:archive` | 手动移动文件 | Archive 做 delta 合并 + 时间戳 + project.md 更新 |

## Red Flags — 这些想法意味着 STOP

| 想法 | 现实 |
|------|------|
| "这太简单了，不需要规约" | 简单变更引起复杂 bug。5 行 proposal.md 省几小时。 |
| "我写完代码再补规约" | 事后规约描述的是你做了什么，不是需要什么。不对齐。 |
| "规约在聊天记录里" | 聊天记录会消失。文件持久。写下来。 |
| "我已经知道要构建什么了" | 知道 ≠ 审核过。规约是共识，不是想法。 |
| "规约拖慢速度" | 预期不对齐的重做更慢。先对齐，再执行。 |
| "这只是个原型" | 原型会变成生产代码。现在就写规约，以后省迁移成本。 |

**以上统统意味着：走 SDD 流程。不要走捷径。**

## 极简起步

如果觉得全流程太复杂，先从这个最小闭环开始：

```
/opsx:propose <名称> → 审核四大件 → 实施 → 测试通过 → /opsx:archive <名称>
```

然后按痛点逐步引入 Superpowers 技能——任务太大加入 `@writing-plans`，频繁回归加入 `@test-driven-development`，口头"搞定"却没真跑测试加入 `@verification-before-completion`。**感受到痛再加药，而非一次全吃。**

## 相关 Skill

| Skill | 类型 | 描述 |
|-------|------|------|
| `sdd-workflow` | Rigid | 分类请求 → 检测阶段 → 路由工具 |
| `sdd-review-specs` | Rigid | 四大件审核 gate |
| `@brainstorming` | Flexible | 苏格拉底式设计对话 |
| `@writing-plans` | Flexible | 任务粒度细化 |
| `@test-driven-development` | Rigid | RED-GREEN-REFACTOR 循环 |
| `@systematic-debugging` | Rigid | 根因调查优先 |
| `@verification-before-completion` | Rigid | 完成前验证证据 |
| `@requesting-code-review` | Rigid | 结构化代码审查 |
