FROM golang:1.16

COPY . /go/src/app

WORKDIR /go/src/app/cmd/server

RUN go build -o server main.go

EXPOSE 9090

CMD ["./server"]
