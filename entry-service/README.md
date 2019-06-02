# Entry Service

To build and push the docker image run:

```sh
bash build-and-push-image.sh {TAG}
```

To build and expose Entry Service with Docker run:

```sh
docker build -t entry-service .
docker run -p 8000:8000 entry-service
```