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