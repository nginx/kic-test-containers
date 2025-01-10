# KIC Test Containers

This repository contains the Dockerfiles and code for the test containers used in the [Kubernetes Ingress Controller](https://github.com/nginx/kubernetes-ingress).

## Building the containers for local testing

For each container, run the following command from the root of the repository:

```console
docker build -t <container-name> -f <container-type>/Dockerfile .
```

Alternatively, you can use the `Makefile` to build the containers. For example, to build the `grpc-server` container,
run the following command from the root of the repository:

```console
make grpc-server
```

## Publishing the containers

The containers are published to the [NGINX GitHub Container Registry](https://github.com/orgs/nginx/packages?repo_name=kic-test-containers)
when a new tag is created in this repository. You can see the packages on the right side of this page.
