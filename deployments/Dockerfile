# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS build
WORKDIR /src
RUN apk add --no-cache git ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /out/telegram-commander ./cmd/telegram-commander

FROM alpine:3.21
RUN apk add --no-cache ca-certificates tzdata
COPY --from=build /out/telegram-commander /usr/local/bin/telegram-commander
USER nobody
ENTRYPOINT ["telegram-commander"]
CMD ["--help"]
