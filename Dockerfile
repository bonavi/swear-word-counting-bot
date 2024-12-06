# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder
RUN apt install ca-certificates
ADD . /swear
WORKDIR /swear/swear
RUN git submodule update --init --recursive
RUN sed -i -e 's|@{build}|12186747288|g' -e 's|@{version}|v1.0.0|g' internal/main.go
RUN go generate ./...
RUN go build -o ../main internal/main.go

FROM ubuntu:latest
WORKDIR /swear-test
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /swear/main .
EXPOSE 8089
USER root:root
CMD ["./main"]
