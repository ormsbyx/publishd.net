# Publishd

Premium reading platform for short stories, essays, and articles.

## Development

```bash
go run main.go
```

Server runs on http://localhost:8080

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