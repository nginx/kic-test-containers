# syntax=docker/dockerfile:1.14
FROM golang:1.24 AS builder

# renovate: datasource=github-releases depName=grpc/grpc-go
ARG GRPC_GO_VERSION=v1.71.1

ADD https://github.com/grpc/grpc-go.git#${GRPC_GO_VERSION} /grpc-go

RUN CGO_ENABLED=0 go build -C /grpc-go/examples/helloworld/greeter_server -o /greeter_server

FROM scratch
COPY --from=builder /greeter_server /
CMD ["/greeter_server"]
