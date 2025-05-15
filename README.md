# Go-lang RESTful API Project Practice From Udemy Course

## Description

A Go-lang RESTful API project practice from [Udemy Course - "Pemrograman Go-Lang : Pemula sampai Mahir"](https://www.udemy.com/course/pemrograman-go-lang-pemula-sampai-mahir/), instructed by [Programmer Zaman Now](https://www.udemy.com/user/eko-kurniawan/).

The project is to develop a create, read, update, and delete (CRUD) category service RESTful API (initially, but will be adding another service in the future). Even though the application is extremely simple and straightforward, this project focused on implementing the best practices for developing Go-lang applications, such as implementing clean architecture, automation testing (unit test, integration test, and E2E test), dependency injection, logging, middleware, etc.

The Clean Architecture is inspired by [https://github.com/khannedy/golang-clean-architecture](https://github.com/khannedy/golang-clean-architecture).

## Tech Stack

- Go-lang: [https://github.com/golang/go](https://github.com/golang/go)
- PostgreSQL (Database): [https://www.postgresql.org/](https://www.postgresql.org/)

## Library

- Golang Migrate (Database Migration) : [https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- Google Wire (Dependencies Injection): [https://github.com/google/wire](https://github.com/google/wire)
- HttpRouter (Router) : [https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- pgxpool (Concurrency-safe connection pool for [pgx](https://github.com/jackc/pgx)) : [https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool](https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool)
- Viper (Configuration) : [https://github.com/spf13/viper](https://github.com/spf13/viper)
- Go Playground Validator (Validation) : [https://github.com/go-playground/validator](https://github.com/go-playground/validator)
- Logrus (Logger) : [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)

### Testing and Mocking

- Testify (Unit testing for Golang) : [https://github.com/stretchr/testify](https://github.com/stretchr/testify)
- pgxmock (pgx driver mock for Golang) : [https://github.com/pashagolub/pgxmock](https://github.com/pashagolub/pgxmock)

## Tool

- Air (Live reload for Go apps): [https://github.com/air-verse/air](https://github.com/air-verse/air)

## API Spec

All API specification is in `api` folder.

## Configuration

The app configuration file must be named `config.yaml`, and an example of its content is in the `config-example.yaml`. For Air (a live reload Go-lang apps tool) configuration, is in `.air.toml`.

## CLI Flags

```bash
Options:
  -configPaths=<path>       Set the config.yaml absolute or relative
                            location file. Default value "./"
```

## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```bash
migrate create -ext sql -dir db/migrations create_examples_table
```

### Run Migration

```bash
migrate -database postgres://user:password@host:port/dbname?query -path db/migration up
```

## Dependency Injection

Enter the directory where the [injector](https://github.com/google/wire/blob/main/docs/guide.md#injectors) is located in your Terminal or CMD. Then execute,

```bash
wire
```

Learn more about [Google Wire](https://github.com/google/wire).

## Run Application

### Run all test

```bash
go test -v -p=1 -count=1 ./internal/... ./test/e2e
```

The testing command will output verbosely (-v), run in sequentially (-p=1), and run without cache (-count=1). The testing can be run concurrently by removing the `-p=1` flag for faster testing execution.

However, it can be a **problem** when several tests must interact with a database or other shared resources. This can cause a race condition, and the tests result will not be reliable, even the fact that the tests should be successful.

### Run web server

```bash
go run ./cmd/web
```

### Run web server (with Air live reload)

```bash
air
```

Learn more about the [Air](https://github.com/air-verse/air) installation and its usages.

### Build

```bash
go build -o ./bin/app ./cmd/web
```

#### Run Binary

```bash
./bin/app # or in Windows, ...\bin\app.exe
```

## License

[MIT](./LICENSE.txt) License (c) 2025-present Syahda Romansyah
