# TODO: Check if port is available.

cd ../auth-service
docker build -t auth-service .
docker run -d -p 8001:8001 auth-service

cd ../entry-service
go test -tags integration

# TODO: Somehow delete running container.