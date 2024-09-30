
## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

## Packages
| Area            | Package Name     | Package Repo                        |
| --------------- | ---------------- | ----------------------------------- |
| Configuration   | Viper            | https://github.com/spf13/viper      |
| Web Framework   | Http/Net         | golang/net/http                     |
| ORM             | SQLC             | https://github.com/sqlc-dev/sqlc    |
| Logging         | Logrus           | golang/log/slog                     |
| Testing         | Testify          | https://github.com/stretchr/testify |
| Templating      | Templ            | https://github.com/a-h/templ        |
| CLI             | Cobra            | https://github.com/spf13/cobra      |
| Migration       | Goose            |                                     |
| Migration       | Golang Migration | https://github.com/golang-migrate/migrate |
| Postgres Driver | PGX              | https://github.com/jackc/pgx        |