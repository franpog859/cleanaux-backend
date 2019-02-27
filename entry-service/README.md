# Entry Service

To build and push the docker image run:

```sh
docker build -t entry-service .
docker tag {IMAGE_ID} franpog859/entry-service:{TAG}
docker push franpog859/entry-service:{TAG}
```

To build and expose Entry Service with Docker run:

```sh
docker build -t entry-service .
docker run -p 8000:8000 entry-service
```