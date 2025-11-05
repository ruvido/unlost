# Unlost - Photo & Video Library

**Less is more!**

## Fine
Self-hosted photo/video library. Auto-organizza per data, album manuali, grid UI.

## Pilastri
- **Backend**: PocketBase (Go framework)
- **Frontend**: Svelte + hash routing
- **Deploy**: Single binary
- **Storage**: `library/username/year/month/day/file.jpg`

## Strategia di Sviluppo
- **Less is more**: codice semplice e chiaro
- **I docs ufficiali sono dio**: consulta SEMPRE docs ufficiali prima di implementare
- **Best practices only**: seguire convenzioni ufficiali
- **Zero ridondanza**: evitare codice complesso e duplicato

## Docs Ufficiali
- PocketBase: https://pocketbase.io/docs/
- Go Framework: https://pocketbase.io/docs/use-as-framework/
- Svelte: https://svelte.dev/docs/svelte/
- Hash Routing: https://svelte.dev/docs/kit/configuration#router

## Comandi
```bash
./pocketbase serve                    # Start dev
cd frontend && npm run build         # Build frontend
go build -o pocketbase main.go       # Build binary
```