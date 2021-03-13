# cron

This service will periodically get live data about all flights by using the OpenSky API and saving them in a PostgreSQL database.

## Running

```
go run main.go
```

```
docker build -t cron . && docker run --rm --detach --env-file .env --name cron-standalone cron
```