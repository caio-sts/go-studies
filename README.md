# I am reading these articles specially from medium to learn go
- https://jeffotoni.medium.com/primeiros-passos-em-go-e1e28b7ff5d3
- 

# Mentioned features
- It is complete to web development, the server part
- It is simple, legible and productive
- just 25 keywords, low learning curve, statically compiled, statically typed, 
multiplatform supportint RISC-V (instruction set for chips), concurrent paradigm,
retrocompatibility

- It is known as container's language

# You write your code, use tidy when you need external packages but if you are
# using just natives you can build when finish using

`go build <codefile>`

# To rename 

`go build -o <new_name> <codefile>`

# To run the compiled

`./<new_name>` or `./<codefile>`

# To run in the server, you just need to send the binary to the container and execute

# Go and Docker

```go
#################################################
# Dockerfile distroless
#################################################
FROM golang:1.15 as builder
WORKDIR /go/src/main
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go
RUN cp main /go/bin/main

############################
# STEP 2 build a small image
############################
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/main /
CMD ["/main"]
```

# This repo has workflows to run in github actions
- run tests and generate coverage report

# Principais bibliotecas em Go

## Biblioteca padrão (mais usadas)
- **fmt** → formatação e impressão
- **os** → manipulação de arquivos, variáveis de ambiente, processos
- **io**, **bufio** → leitura/escrita de streams
- **net/http** → criar servidores e clientes HTTP
- **encoding/json**, **encoding/xml**, **encoding/csv** → serialização de dados
- **time** → datas, horários, timers
- **sync**, **sync/atomic** → sincronização de goroutines
- **context** → controle de cancelamento/timeout em goroutines e chamadas externas
- **database/sql** → interface genérica para bancos relacionais
- **testing** → framework nativo para testes

## Bibliotecas da comunidade (mais comuns)
### Web e APIs
- [gin-gonic/gin](https://github.com/gin-gonic/gin) → framework HTTP rápido e popular
- [gorilla/mux](https://github.com/gorilla/mux) → roteador HTTP flexível
- [labstack/echo](https://github.com/labstack/echo) → framework minimalista para APIs

### Banco de dados
- [gorm](https://gorm.io) → ORM para Go
- [jmoiron/sqlx](https://github.com/jmoiron/sqlx) → extensão de `database/sql` com helpers

### Configuração
- [spf13/viper](https://github.com/spf13/viper) → gerenciamento de configs (JSON, YAML, ENV)

### CLI e utilitários
- [spf13/cobra](https://github.com/spf13/cobra) → criação de CLIs
- [urfave/cli](https://github.com/urfave/cli) → alternativa leve para CLIs

### Testes e mocks
- [stretchr/testify](https://github.com/stretchr/testify) → asserts, mocks
- [golang/mock](https://github.com/golang/mock) → geração de mocks para interfaces

### Logs
- [sirupsen/logrus](https://github.com/sirupsen/logrus) → logging estruturado
- [uber-go/zap](https://github.com/uber-go/zap) → logging de alta performance

### Concorrência e mensagens
- [google.golang.org/protobuf](https://pkg.go.dev/google.golang.org/protobuf) → suporte a Protobuf
- [nats-io/nats.go](https://github.com/nats-io/nats.go) → pub/sub distribuído
