FROM golang:1.10-alpine as builder

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH   

RUN mkdir -p $GOPATH/src/app 
ADD . $GOPATH/src/app
WORKDIR $GOPATH/src/app 

RUN go build -o app
RUN go test -c ./...

FROM alpine:3.8
COPY --from=builder /go/src/app/app .
ENTRYPOINT ./app

EXPOSE 8000