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

