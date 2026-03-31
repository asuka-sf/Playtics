# Playtics

## Overview

Playtics is a game stats API that tracks player scores, match results, and leaderboard rankings — built with Go, PostgreSQL, and Clean Architecture.

## Tech Stack

- **Language**: Go
- **Architecture**: Clean Architecture
- **Database**: PostgreSQL
- **Schema Management**: Atlas (ariga.io)
- **Infrastructure**: AWS
- **IaC**: Terraform

## Code Style

- Follow standard Go conventions (`gofmt`, `goimports`)
- Error handling: always handle errors explicitly, no `_` for error returns
- Use dependency injection via constructor functions (e.g. `NewXxx()`)
- Interface definitions belong in the domain layer, not infrastructure
- Keep handler thin: validation and response only, no business logic
- Use context (`ctx`) as the first argument for all functions that access DB

## Architecture

This project follows Clean Architecture principles, separating concerns into distinct layers:
- **Handler**: HTTP request/response handling
- **Usecase**: Business logic
- **Domain**: Entities and repository interfaces
- **Infrastructure**: Database implementation
Dependencies flow inward:
```
Handler → Usecase → Domain ← Infrastructure
```

### Directory Structure

- `cmd/` - Entry point
- `internal/handler/` - HTTP request/response handling
- `internal/usecase/` - Business logic
- `internal/domain/` - Entities and repository interfaces
- `internal/infrastructure/postgres/` - DB implementation
- `internal/config/` - Configuration
- `internal/registry/` - Dependency injection
- `terraform/` - AWS infrastructure

## API

See `docs/api.md` for full request/response documentation.

## Development Rules

### Commit Message Format

- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation
- `refactor:` - Refactoring
- `test:` - Tests
> Commit Messages should have gitmoji like ✨, 🐛 and ♻️
> "♻️ refactor: API response"

## Review Guidelines
- Always check that dependency directions follow Clean Architecture
- Domain layer must not depend on infrastructure layer
- All repository interfaces must be defined in domain layer

## Test plan
- [ ] Base branch is `develop` (not `main`)
- [ ] `go build` passes
- [ ] `go test` passes
- [ ] No hardcoded credentials or secrets
- [ ] API documentation updated (`docs/api.md`)