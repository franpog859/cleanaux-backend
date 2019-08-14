FROM golang:1.10-alpine as builder

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/github.com/franpog859/cleanaux-backend/auth-service
ADD . $GOPATH/src/github.com/franpog859/cleanaux-backend/auth-service
WORKDIR $GOPATH/src/github.com/franpog859/cleanaux-backend/auth-service

RUN go build -o app ./cmd/authservice/...
RUN go test ./...

FROM alpine:3.8
COPY --from=builder /go/src/github.com/franpog859/cleanaux-backend/auth-service/app .
ENTRYPOINT ./app

EXPOSE 8000
EXPOSE 8001