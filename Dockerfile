# 1. Stage for build the go project
FROM golang:1.24 AS build-stage
WORKDIR /app

ENV GOPROXY="https://goproxy.io,direct"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /learnyscape-backend-mono/app

# 2. Stage for create a lightweight runtime container
FROM alpine:3.21.3 AS build-release-stage
WORKDIR /

COPY --from=build-stage /learnyscape-backend-mono/app /learnyscape-backend-mono/app
COPY ./.env .

EXPOSE 8080

ENTRYPOINT [ "/learnyscape-backend-mono/app", "go run ./cmd/api"]
