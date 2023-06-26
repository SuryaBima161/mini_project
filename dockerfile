# FROM golang:1.18-alpine3.16 as builder

# WORKDIR /app
# COPY . .

# RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -o main.app .

# FROM alpine:latest

# COPY --from=builder /app/main.app .

# CMD ["./main.app"]
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /docker-api

EXPOSE 8080

CMD ["/docker-api"]

