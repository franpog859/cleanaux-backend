# Cleanaux Backend

**WORK IN PROGRESS** - It's just my Go, Docker and Kubernetes playground for now. After some work on authorization, monitoring and acceptance tests it will become a cloud backend for the Cleanaux app. It's public because of specific database initialization which requires downloading files.

Cleanaux provides a list of things which should be cleaned regularly. If something wasn't cleaned for a long time it appears with a warning color. After cleaning whole list is updated.

Whole backend is deployed on Google Cloud Platform with Google Kubernetes Engine.

## Content Service

Go microservice exposing REST API via Ingress. It uses MySQL database to store data. Database runs on separate pod and is initialized with an init container.

## Structure

- `*-service/` directory contains specific service files
- `kube/` directory contains Kubernetes files for cloud deployment
- `scripts/` directory contains installation scripts
- `test/` directory contains some testing tips

## Development

To build Docker images from local repository and push them to the Docker Hub run:

```bash
bash scripts/build-and-push-images.sh {TAG}
```

To mock an interface go to its directory and run:

```bash
$GOPATH/bin/mockery -name={INTERFACE_NAME}
```

## Production

Before installing the Cleanaux Backend on your cluster edit the default `jwtkey` value in the `kube/auth-service/secret.yaml` file.

To add a new user run MongoDB client and run script provided in `auth-service/init/db-test` with your values instead of default ones. Remember to put here a base64 encoded password. You can encode it running:

```bash
echo -n '{PASSWORD}' | base64
```

## To do

- create Swagger API description for both services
- create end to end test app
