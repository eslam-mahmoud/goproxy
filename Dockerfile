FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN go build -o goproxy

EXPOSE 8081

CMD ["./goproxy"]