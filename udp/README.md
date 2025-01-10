# UDP Server

A Go server that accepts UDP requests and responds with the local address of the connection.

## Description

If the server is run inside a Docker container, the local address is the IP of the docker container. This is useful
for distinguishing between instances of Docker containers. This server is used by the python tests in the
[load balancing tests](../suite/test_transport_server_udp_load_balance.py).

## Config

The default port the server listens to is `3334`. The server takes a single argument, `port`, to allow the port to be
overridden.

## Making changes

- Test the change:
  - Use the minikube registry ```$ eval $(minikube docker-env)```
  - Build the docker image from the root of the project ```docker build -t udp-server .```
  - Update the [service yaml](https://github.com/nginx/kubernetes-ingress/blob/main/tests/data/transport-server-udp-load-balance/standard/service_deployment.yaml)
  to use the local version ```-> imagePullPolicy: Never```
  - Test the changes
- Publish the upd-server change
  - Open a PR with the changes
  - Once the PR is approved merge it to the main branch
  - Create a new tag for the commit, this will start the pipeline and publish the new udp-server image
