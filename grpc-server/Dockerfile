FROM golang as builder

COPY ./protos /go/src/github.com/jpreese/grpc-catfact/grpc-server/protos
COPY ./server.go .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/server


FROM scratch
COPY --from=builder /bin/server .

ENTRYPOINT ["./server"]
