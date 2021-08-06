FROM golang:1.14.4 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app/src
RUN go build -o main .
CMD ["/app/src/main"]