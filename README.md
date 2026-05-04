# go-movies-common

Shared Go library for the go-movies microservices ecosystem. Provides common utilities for configuration, database connectivity, HTTP middleware, observability, and Kafka event types.

## Packages

### pkg/config
Environment variable loading utilities.

```go
import "github.com/charbelhanna96/go-movies-common/pkg/config"

port := config.GetEnv("PORT", "8080")
maxConns := config.GetEnvInt("DB_MAX_CONNS", 25)
brokers := config.GetEnvList("KAFKA_BROKERS", []string{"kafka:9092"})
```

### pkg/db
PostgreSQL connection pool setup.

```go
import "github.com/charbelhanna96/go-movies-common/pkg/db"

database, err := db.Connect(db.DatabaseConfig{
    Host:     "localhost",
    Port:     "5432",
    Name:     "mydb",
    User:     "postgres",
    Password: "password",
})
```

### pkg/web
HTTP response helpers.

```go
import "github.com/charbelhanna96/go-movies-common/pkg/web"

web.JSON(w, http.StatusOK, payload)
web.Error(w, http.StatusBadRequest, "invalid input")
```

### pkg/middleware
HTTP middleware for CORS, metrics, and tracing.

```go
import "github.com/charbelhanna96/go-movies-common/pkg/middleware"

handler := middleware.CORS(allowedOrigins, next)
handler := middleware.Tracing("my-service", next)
handler := middleware.Metrics(requestCount, requestDuration, next)
```

### pkg/tracing
OpenTelemetry setup with Jaeger export.

```go
import "github.com/charbelhanna96/go-movies-common/pkg/tracing"

shutdown, err := tracing.Setup(ctx, "jaeger:4318", "my-service")
defer shutdown(ctx)
```

### pkg/kafka
Shared Kafka event types and topic constants.

```go
import "github.com/charbelhanna96/go-movies-common/pkg/kafka"

event := kafka.SearchEvent{
    Filters:      kafka.SearchFilters{GenreIDs: []int{1, 2}},
    ResultsCount: 10,
    Timestamp:    time.Now().UTC(),
}
```

## Shared Infrastructure

A `docker-compose.yml` is provided with shared infrastructure services (Kafka, Jaeger) that can be included in service-level compose files.

```yaml
include:
  - path: ../go-movies-common/docker-compose.yml
```

## Services Using This Library

- [go-movies-api](https://github.com/charbelhanna96/go-movies-api) — movie search REST API
- [go-movies-analytics](https://github.com/charbelhanna96/go-movies-analytics) — search analytics service

## Testing

```bash
go test ./... -v
```