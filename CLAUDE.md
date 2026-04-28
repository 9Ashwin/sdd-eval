# OpenSpec + Superpowers Guide

Evaluation and demonstration project for the OpenSpec + Superpowers SDD workflow. Multi-language test apps for skill validation.

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
