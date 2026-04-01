# API Documentation

## Phase 1

| Method | URL | Description | Note |
|--------|-----|-------------|------|
| POST | `/players` | Create a player | |
| POST | `/matches` | Create a match | |
| POST | `/matches/{id}/results` | Create a match result | |
| GET | `/players/{id}` | Get a player | |
| GET | `/leaderboard` | Get leaderboard ranking | Sorted by total score |

## Request & Response

### Error Response

All endpoints return the following format on error:

```json
{
    "success": "NG",
    "message": "error description"
}
```

| Status Code | Description |
|-------------|-------------|
| 400 | Bad Request - Validation error or invalid input |
| 404 | Not Found - Resource does not exist |
| 409 | Conflict - Resource already exists (e.g., duplicate email) |
| 500 | Internal Server Error |

---

### POST /players

**Request**
```json
{
    "name": "Alice",
    "email": "alice@example.com",
    "image_url": "image.jpeg"
}
```

**Response**
```json
{
    "success": "OK",
    "message": "Created player successfully",
    "data": {
        "id": "uuid",
        "name": "Alice",
        "email": "alice@example.com",
        "image_url": "image.jpeg",
        "created_at": "2026-03-25T10:00:00Z",
        "updated_at": "2026-03-25T10:00:00Z"
    }
}
```

**Error Responses**

| Status Code | Message |
|-------------|---------|
| 400 | name is required, email is required |
| 409 | email already exists |
| 500 | internal server error |

### POST /matches

**Request**
```json
{
    "duration_seconds": 300
}
```

**Response**
```json
{
    "success": "OK",
    "message": "Created match successfully",
    "data": {
        "id": "uuid",
        "duration_seconds": 300,
        "created_at": "2026-03-25T10:00:00Z"
    }
}
```

### POST /matches/{id}/results
**Request**
```json
{
    "player_id": "uuid",
    "kill_count": 10,
    "death_count": 5,
    "score": 100
}
```

**Response**
```json
{
    "success": "OK",
    "message": "Created result successfully",
    "data": {
        "player_id": "uuid",
        "match_id": "uuid",
        "kill_count": 10,
        "death_count": 5,
        "score": 100,
        "created_at": "2026-03-25T10:00:00Z",
        "updated_at": "2026-03-25T10:00:00Z"
    }
}
```

### GET /players/{id}

**Error Responses**

| Status Code | Message |
|-------------|---------|
| 400 | invalid player id |
| 404 | player not found |
| 500 | internal server error |

**Response**
```json
{
    "success": "OK",
    "message": "Fetched player successfully",
    "data": {
        "id": "uuid",
        "name": "Alice",
        "email": "alice@example.com",
        "image_url": "image.jpeg",
        "created_at": "2026-03-25T10:00:00Z",
        "updated_at": "2026-03-25T10:00:00Z"
    }
}
```

### GET /leaderboard

**Query Parameters**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| period | string | No | `all` (default), `daily`, `weekly`, `monthly` |
| limit | integer | No | Max rows to return (default: 100) |
| offset | integer | No | Pagination offset (default: 0) |

**Response**
```json
{
    "success": "OK",
    "message": "Fetched leaderboard successfully",
    "data": [
        {
            "rank": 1,
            "player_id": "uuid",
            "player_name": "Alice",
            "total_score": 15800,
            "total_kills": 310,
            "total_deaths": 120            
        },
        {
            "rank": 2,
            "player_id": "uuid",
            "player_name": "Bob",
            "total_score": 10200,
            "total_kills": 250,
            "total_deaths": 200
        }
    ]
}
```