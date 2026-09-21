# GORM template

Provisioned from [`Qode-Platform/fleet-template-v1`](https://github.com/Qode-Platform/fleet-template-v1) - the fleet
lifecycle contract with a GORM starter on top.

## Verified

Built and tested locally on Go 1.23.4 (toolchain auto-upgraded to 1.25):
`go build ./...` and `go test ./...` both pass.

## Fleet lifecycle

| step | command |
|---|---|
| install | `go mod download` |
| build | `go build -o ./.bin/app ./cmd/app` |
| start | `(none - not a service)` |

## Notes

- NOT A SERVICE: START_CMD is empty by design.
- The driver is glebarez/sqlite (pure Go) because the image builds with CGO_ENABLED=0 into distroless, where cgo-based mattn/go-sqlite3 cannot link.
- DB_DSN defaults to file::memory: so the read-only distroless image can still run it.
