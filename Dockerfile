FROM golang:1.24.4-alpine AS builder

COPY . /github.com/AndreiShkolnyi/go-auth/source/
WORKDIR /github.com/AndreiShkolnyi/go-auth/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/AndreiShkolnyi/go-auth/source/bin/crud_server .

CMD ["./crud_server"]