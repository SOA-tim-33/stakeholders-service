FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 8000
ENTRYPOINT ["go", "run", "main.go"]