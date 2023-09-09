FROM golang:alpine

RUN apk update &&  \
    apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /main

EXPOSE 8080

CMD ["/main"]