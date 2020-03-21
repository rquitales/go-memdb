# Copyright (c) 2020 Ramon Quitales
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

FROM golang:1.14.1-alpine AS builder
WORKDIR /go/memdb/
COPY . /go/memdb/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o go-memdb .

FROM alpine:latest
WORKDIR /memdb/
COPY --from=builder /go/memdb/go-memdb /memdb/go-memdb
CMD ["./go-memdb"]