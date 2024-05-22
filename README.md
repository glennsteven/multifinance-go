## Framework & Library

- GorillaMux (HTTP Framework) : https://pkg.go.dev/github.com/gorilla/mux
- Viper (Configuration) : https://github.com/spf13/viper
- Validator : https://github.com/go-playground/validator
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Logrus (Logger) : https://github.com/sirupsen/logrus

## Database Migration
All database migration is in `db/migrations` folder.

## Configuration
All configuration is in `config.json` file.

## instructions On How To Create The Data Are in The Route File
All route is in `internal/route/route.go`

## Download Dependency
```shell
go mod download
go mod tidy
```

### Run Application
```shell
go run main.go
```

### Create Migrations
```shell
migrate create -ext sql -dir db/migrations {table_name}
```

### Database Migration
migration up
```shell
migrate -database "{driver}://{username}:{password}@tcp({host}:{port})/{database_name}" -path db/migrations up

```

migration down
```shell
migrate -database "{driver}://{username}:{password}@tcp({host}:{port})/{database_name}" -path db/migrations down
```