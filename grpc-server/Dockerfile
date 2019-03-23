# Copy over the required dependencies and build the binary
FROM golang

COPY ./protos /go/src/github.com/jpreese/grpc-catfact/grpc-server/protos
COPY ./server.go .

RUN go get -d -v
RUN go build -o /go/bin/server

ENTRYPOINT ["/go/bin/server"]
