# syntax=docker/dockerfile:1.19
FROM golang:1.25-alpine AS builder

WORKDIR /go/src/github.com/nginx/kic-test-containers/
COPY tcp/main.go ./tcp/
COPY go.mod utils.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./tcp/

FROM scratch AS goreleaser
ARG TARGETARCH

COPY dist/tcp_linux_${TARGETARCH}*/main /app/main

ENTRYPOINT ["/app/main"]

FROM scratch

COPY --from=builder /main /app/

ENTRYPOINT ["/app/main"]
