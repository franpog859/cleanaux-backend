# TODO: Check if port is available.

cd ../auth-service
docker build -t auth-service .
CONTAINER="$(docker run -d -p 8001:8001 auth-service)"

cd ../entry-service
go test -tags integration

docker kill ${CONTAINER}
