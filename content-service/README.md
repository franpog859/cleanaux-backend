# Content Service

To build and push the docker image run:

```sh
docker build -t content-service .
docker tag {IMAGE_ID} franpog859/content-service:{TAG}
docker push franpog859/content-service:{TAG}
```

To build and expose Content Service with Docker run:

```sh
docker build -t content-service .
docker run -p 8000:8000 content-service
```