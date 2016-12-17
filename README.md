## Shell

```bash
make shell
```

## Build

```bash
make build
```

## Run

```bash
LOG_LEVEL=debug make run
```

## Database

```bash
make db
```

## Run database migration

```bash
make db-up
```

## Create a new migration

```bash
goose -dir db/migrations create AddSomeColumns sql
```
