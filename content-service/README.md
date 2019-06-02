# Content Service

To build and push the docker image run:

```sh
bash build-and-push-image.sh {TAG}
```

To build and expose Content Service with Docker run:

```sh
docker build -t content-service .
docker run -p 8000:8000 content-service
```