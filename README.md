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

### Web Interface

#### Reader-Facing Pages
- `GET /` - Homepage with featured stories
- `GET /stories` - Browse all published stories
- `GET /stories/:id` - Read individual story (Kindle-like experience)

#### Admin Interface
- `GET /admin` - Story management interface

### API Endpoints

#### Health & Info
- `GET /health` - Health check with database status

#### JSON API Routes (`/api/v1`)
- `GET /api/v1/stories` - List all stories (published only)
- `GET /api/v1/stories/:id` - Get specific story
- `POST /api/v1/stories` - Create new story
- `PUT /api/v1/stories/:id` - Update story
- `DELETE /api/v1/stories/:id` - Delete story
- `POST /api/v1/stories/:id/publish` - Publish story

#### Legacy JSON Routes
- `GET /api/stories` - List published stories (JSON)
- `GET /api/stories/:id` - Get specific story (JSON)

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