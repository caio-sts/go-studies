#################################################
# Dockerfile distroless
#################################################
FROM golang:1.25 AS builder
WORKDIR /go/src/main
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main 10.simple_and_small_applications.go
RUN cp main /go/bin/main

############################
# STEP 2 build a small image
############################
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/main /
CMD ["/main"]