# Auth Service

To build and push the docker image run:

```sh
docker build -t auth-service .
docker tag {IMAGE_ID} franpog859/auth-service:{TAG}
docker push franpog859/auth-service:{TAG}
```

To build and expose Auth Service with Docker run:

```sh
docker build -t auth-service .
docker run -p 8000:8000 auth-service
```