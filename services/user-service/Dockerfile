# build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o user-service main.go 


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-service .

EXPOSE 8080

CMD ["./user-service"] 