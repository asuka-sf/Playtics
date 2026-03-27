# Playtics

## Overview
Playtics is a game stats API that tracks player scores, match results, and leaderboard rankings — built with Go, PostgreSQL, and Clean Architecture.

## Tech Stack

- **Language**: Go
- **Architecture**: Clean Architecture
- **Database**: PostgreSQL
- **Schema Management**: Atlas (ariga.io)
- **Code Generation**: sqlc
- **Infrastructure**: AWS
- **IaC**: Terraform

## Architecture

This project follows Clean Architecture principles, separating concerns into distinct layers:
- **Handler**: HTTP request/response handling
- **Usecase**: Business logic
- **Domain**: Entities and repository interfaces
- **Infrastructure**: Database implementation

## Features

### Phase 1: Leaderboard API

- Player registration
- Score submission
- Ranking retrieval (global / by period)

### Phase 2: Player Stats API (WIP)

- Match history recording
- Win rate, K/D ratio, average score aggregation

## Development

### Prerequisites

- Go
- Docker
- [Atlas](https://atlasgo.io/)
- [sqlc](https://sqlc.dev/)
- [golangci-lint](https://golangci-lint.run/)

### Commands

| Command | Description |
|---------|-------------|
| `make db-up` | Start PostgreSQL (Docker) |
| `make db-down` | Stop PostgreSQL |
| `make migrate-diff name=xxx` | Generate migration diff with Atlas |
| `make migrate-apply` | Apply migrations |
| `make sqlc` | Generate Go code from SQL queries (sqlc) |
| `make run` | Run the application |
| `make lint` | Run linter |
| `make test` | Run tests |