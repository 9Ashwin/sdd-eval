# SDD Eval

SDD skill evaluation workspace — multi-language test apps for skill validation with behavior-based evidence collection.

## Project Structure

```
evaluation/
  go/    — Go task manager API (stdlib net/http)
  py/    — Python (future)
```

## Commands

- `cd evaluation/go && go test ./...` — Run all tests
- `cd evaluation/go && go build ./cmd/server/` — Build
- `cd evaluation/go && go run ./cmd/server/` — Start server on :3000

## Development Workflow

| Change Type | Workflow |
|-------------|----------|
| One-line fix, comment | Make change directly, verify with tests |
| New feature / behavior change | `sdd-workflow` |
| Bug with unclear cause | `@systematic-debugging` |

`sdd-workflow` handles classification, routing, and all phase transitions internally.

## Reference

- [SDD Workflow Guide (中文)](docs/guide-zh.md) — full methodology, tool matrix, red flags
