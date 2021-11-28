FROM golang:latest

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

RUN export GO111MODULE=on

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o handler ./cmd/main.go

CMD ["./handler"]