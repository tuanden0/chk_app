FROM golang:1.18.1-alpine3.15 AS base

RUN apk add build-base

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./chk ./cmd/chk_apis/v1/main.go

FROM alpine:latest

COPY --from=base /app/chk /app/chk

WORKDIR /app

RUN mkdir upload

EXPOSE 8080

CMD ["./chk"]