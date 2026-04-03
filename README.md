# cicd-ref

`cicd-ref` is a small Go HTTP service used as a reference project for local development, container builds, and CI/CD practice.

The application starts a web server, reads configuration from environment variables, and exposes a couple of simple JSON endpoints that are easy to validate in pipelines and deployment environments.

## Features

- Lightweight Go API with no external runtime dependencies
- Health endpoint for readiness/smoke checks
- Message endpoint driven by environment configuration
- Multi-stage Docker build
- Kubernetes manifests for `dev` and `prod`
- Handler tests for core API behavior

## API

The server currently registers these routes:

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/health` | Returns service health status |
| `GET` | `/message` | Returns the configured application message |

### Example responses

`GET /health`

```json
{
  "status": "ok"
}
```

`GET /message`

```json
{
  "message": "secure pipeline demo"
}
```

## Configuration

Configuration is loaded from environment variables in `internal/config/config.go`.

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `8080` | Port the HTTP server listens on |
| `APP_MESSAGE` | `secure pipeline demo` | Message returned by `GET /message` |

## Project structure

```text
cmd/api/main.go                Application entrypoint
internal/config/config.go      Environment-based configuration
internal/handlers/handlers.go  HTTP handlers
internal/server/server.go      Route registration and server wiring
tests/handlers_test.go         API handler tests
deploy/dev/deployment.yaml     Development Kubernetes manifest
deploy/prod/deployment.yaml    Production Kubernetes manifest
Dockerfile                     Multi-stage container build
```

## Local development

### Prerequisites

- Go `1.25` or newer (matches `go.mod`)
- `make` (optional, for convenience targets)
- Docker (optional, for container builds)

### Common commands

```bash
make tidy
make test
make build
make run
```

### Direct Go commands

```bash
go mod tidy
go test ./...
go build -o bin/api ./cmd/api
go run ./cmd/api
```

### Run with custom configuration

```bash
PORT=9090 APP_MESSAGE="hello pipeline" go run ./cmd/api
```

Then verify the service:

```bash
curl -s http://localhost:9090/health
curl -s http://localhost:9090/message
```

## Testing

The test suite currently covers the JSON responses returned by the handler layer.

Run tests with:

```bash
go test ./...
```

To force a fresh run without cached results:

```bash
go test ./... -count=1
```

## Docker

The project includes a multi-stage `Dockerfile` that builds the Go binary in an Alpine-based Go image and runs it in a smaller Alpine runtime image.

### Build the image

```bash
docker build -t cicd-ref:local .
```

### Run the container

```bash
docker run --rm -p 8080:8080 cicd-ref:local
```

With custom message:

```bash
docker run --rm -p 8080:8080 -e APP_MESSAGE="hello from docker" cicd-ref:local
```

## Kubernetes manifests

Deployment manifests are included under `deploy/`:

- `deploy/dev/deployment.yaml`
- `deploy/prod/deployment.yaml`

The development manifest runs a single replica and sets `APP_MESSAGE` to `running in dev`.

The production manifest runs two replicas and sets `APP_MESSAGE` to `running in prod`.

Example apply command:

```bash
kubectl apply -f deploy/dev/deployment.yaml
```

## Notes

- The Go module path is `CICDRef`.
- The server entrypoint is `cmd/api/main.go`.
- The HTTP routes are registered in `internal/server/server.go`.
