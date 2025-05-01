# renovate: datasource=github-tags depName=golangci/golangci-lint
GOLANGCI_LINT_VERSION = v2.1.5

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run --fix

build-docker-tcp:
	docker build -t tcp-server -f tcp/Dockerfile .

build-docker-udp:
	docker build -t udp-server -f udp/Dockerfile .

build-docker-grpc:
	docker build -t grpc-server -f grpc/Dockerfile .

build-docker-nap-dos:
	docker build -t nap-dos-server -f nap-dos/Dockerfile .

run-docker-tcp:
	docker run -p 3333:3333 -it tcp-server

run-docker-udp:
	docker run -p 3334:3334 -it udp-server

run-docker-grpc:
	docker run -p 50051:50051 -it grpc-server

run-docker-nap-dos:
	docker run -p 8000:8000 -it nap-dos-server
