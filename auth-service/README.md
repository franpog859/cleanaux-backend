# Auth Service

To build and push the docker image run:

```sh
bash build-and-push-image.sh {TAG}
```

To build and expose Auth Service with Docker run:

```sh
docker build -t auth-service .
docker run -p 8000:8000 -p 8001:8001 auth-service
```
