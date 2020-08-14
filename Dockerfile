FROM golang:alpine

WORKDIR /superapi

COPY go.mod go.sum ./

RUN go mod download

COPY . .
