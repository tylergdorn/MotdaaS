FROM golang:latest as build

WORKDIR /go/src/github.com/tylergdorn/MotdaaS

RUN go get -u "github.com/gobuffalo/packr/v2/packr2"
RUN go get -u "google.golang.org/grpc"
RUN go get -u "github.com/golang/protobuf/proto"

COPY ./motd ./motd
COPY ./server ./server

RUN (cd server; packr2 clean)
RUN (cd server; packr2)
RUN mkdir /app
RUN (cd server; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .)

FROM alpine:latest
COPY --from=build /app /app
EXPOSE 7777

CMD ["/app/server"]