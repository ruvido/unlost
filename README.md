# Unlost

Self-hosted photo and video library. Less is more.

## Features

- Auto-organize by date
- Manual albums
- User isolation
- Mobile-first grid UI

## Tech Stack

- **Backend**: PocketBase (Go framework)
- **Frontend**: Svelte + Vite
- **Storage**: Filesystem organized by `library/username/year/month/day/`

## Quick Start

```bash
# Build
go build -buildvcs=false -o pocketbase
cd frontend && npm install && npm run build && cd ..

# Run
./pocketbase serve
```

Server runs on http://localhost:8090

## Development

```bash
# Backend
./pocketbase serve

# Frontend (dev mode)
cd frontend && npm run dev
```

## Architecture

- Single binary deployment
- Hash-based deduplication
- Two-phase scanner (filesystem→DB, DB→filesystem)
- Owner-based security rules

## License

MIT
