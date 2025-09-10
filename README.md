# Publishd

Premium reading platform for short stories, essays, and articles.

## Development

### Prerequisites
- Go 1.25+
- PostgreSQL database

### Setup
1. Copy environment variables:
   ```bash
   cp .env.example .env
   ```

2. Update `.env` with your database credentials

3. Start the server:
   ```bash
   go run main.go
   ```

Server runs on http://localhost:8080

### API Endpoints
- `GET /` - Welcome message
- `GET /stories` - List all stories  
- `GET /stories/:id` - Get specific story
- `GET /health` - Health check with database status

## Project Structure

```
publishd.net/
├── main.go                 # Application entry point
├── cmd/server/            # Server-specific code
├── internal/              # Private application code
│   ├── handlers/          # HTTP handlers
│   ├── models/           # Data models
│   └── database/         # Database connection and queries
├── web/                  # Web assets
│   ├── templates/        # HTML templates
│   └── static/          # CSS, JS, images
└── config/              # Configuration files
```

## Tech Stack

- **Backend:** Go
- **Database:** PostgreSQL  
- **Frontend:** Vue 3
- **Hosting:** Render