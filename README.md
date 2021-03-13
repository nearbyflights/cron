# cron

This service will periodically get live data about all flights by using the OpenSky API and saving them in a PostgreSQL database.

## Running

```
go run main.go
```

```
docker build -t cron . && docker run --rm --detach --env-file .env --name cron-standalone cron
```

## Create new release

For creating a new release just add and push a new tag, GoReleaser will automatically create a new GitHub release and add all artifacts to it.

```
git tag -a v1.0 -m "First release"
git push origin --tags
```