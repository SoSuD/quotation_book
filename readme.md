Тестовое задание.

В `.env` задаются переменные:

* `DATABASE_URL`
* `SERVER_PORT`

Миграции применяются с помощью golang-migrate.

Установка golang-migrate:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Применение миграций:

```bash
migrate -path migrations -database "postgres://brandscout:brandscout@localhost:5432/brandscout?sslmode=disable" up
```
