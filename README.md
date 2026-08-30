# go-studies

Learning Go by building a small HTTP service properly — with tests, CI and a container image —
rather than by collecting snippets.

## What's here

- **`api/ping.go`** — an HTTP handler returning `{"msg":"hello_world"}` with a 200
- **`tests/`** — `server_responding_test.go` and `10.simple_api_test.go`
- **`.github/workflows/ci-go.yml`** — on every PR to `main`: `go build ./...`, `go vet ./...`,
  then `go test` with atomic coverage, a coverage summary and an HTML report uploaded as an
  artifact
- **`Dockerfile`** — two-stage build: compile with `CGO_ENABLED=0` on `golang:1.25`, then copy
  the single binary into `gcr.io/distroless/base`

Go 1.25.

## Running it

```bash
go build ./...
go vet ./...
go test -cover ./...
```

## Note

The Dockerfile still builds `10.simple_and_small_applications.go`, a file that is no longer in
the repository — point it at the current entrypoint before building the image.

## Reference

Notes were taken while working through
[jeffotoni's Go series](https://jeffotoni.medium.com/primeiros-passos-em-go-e1e28b7ff5d3).
