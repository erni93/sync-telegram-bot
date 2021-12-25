# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17.5-alpine3.15 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /syncTelegramBot

##
## Deploy
##
FROM alpine:3.15.0

WORKDIR /

COPY --from=build /syncTelegramBot /syncTelegramBot

ENTRYPOINT ["/syncTelegramBot"]