## Prisma Go Template

An api using prisma as the database orm and gin as the web framework.

### Running the database and server

```bash
go mod download
```

Generate the prisma client

```bash
go run github.com/steebchen/prisma-client-go db push
```

Run the server

```bash
go run . # or go run main.go
```