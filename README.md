# OpenSpec + Superpowers Guide

结合 [OpenSpec](https://github.com/Fission-AI/OpenSpec)（规约层）与 [Superpowers](https://github.com/9Ashwin/superpowers)（执行纪律层）的 SDD（Spec-Driven Development）工作流演示与评估项目。

## 项目结构

```
├── CLAUDE.md              # 项目配置，路由新功能到 sdd-workflow
├── AGENTS.md → CLAUDE.md  # 软链接，无需手动同步
├── .claude/skills/
│   ├── sdd-workflow/      # SDD 路由器 skill
│   └── sdd-review-specs/  # 规约审核 gate skill
├── evaluation/            # 评估目标应用（按语言组织）
│   ├── go/                # Go 评估应用
│   │   ├── cmd/server/    # HTTP 入口
│   │   ├── internal/      # store + handler + tests
│   │   └── go.mod
│   └── py/                # Python（将来）
├── openspec/              # SDD 规约存储
│   ├── specs/             # 永久规约库
│   └── changes/           # 进行中的变更
└── docs/
    └── superpowers/
        ├── specs/         # 设计规约
        └── plans/         # 实施计划
```

## 安装 Skill

```bash
npx skills add https://github.com/9Ashwin/on-my-sdd.git --skill sdd-workflow
npx skills add https://github.com/9Ashwin/on-my-sdd.git --skill sdd-review-specs
```

Skill 源码仓库：[on-my-sdd](https://github.com/9Ashwin/on-my-sdd)

## 两个核心 Skill

本项目只定义了两个 SDD skill：

| Skill | 职责 | 何时触发 |
|-------|------|---------|
| `sdd-workflow` | 分类请求、检测项目阶段、路由到正确的规约或执行工具 | 任何超出单行修复的开发任务 |
| `sdd-review-specs` | 审查 AI 生成的四大件（proposal/specs/design/tasks），设定审核深度（Tier 0/1/2） | `/opsx:propose` 生成规约后 |

`sdd-workflow` 内部处理完整的 SDD 管道，CLAUDE.md 只需一行引用。

## 快速开始

### 运行 Go 评估应用

```bash
cd evaluation/go
go test ./...           # 运行测试
go run ./cmd/server/    # 启动服务 :3000
```

### API

```
GET    /tasks              # 列出所有任务
POST   /tasks              # 创建任务 {"title": "..."}
DELETE /tasks/{id}         # 删除任务
```

### 触发 SDD 流程

直接在对话中描述需求，CLAUDE.md 会自动路由到 `sdd-workflow`，例如：

> "帮我加一个 PATCH /tasks/{id} 接口来切换任务完成状态"

## 添加新语言评估

1. 在 `evaluation/` 下创建语言目录，如 `evaluation/py/`
2. 实现一个最小可运行应用（增删查改即可）
3. 在 `docs/superpowers/specs/` 写评估设计规约
4. 在 `docs/superpowers/plans/` 写实施计划
5. 按 RED → GREEN → REFACTOR 执行评估

## 相关资源

- [SDD 工作流指南（中文）](docs/guide-zh.md)
- [OpenSpec](https://github.com/Fission-AI/OpenSpec) — 规约管理
- [Superpowers](https://github.com/9Ashwin/superpowers) — 执行纪律
