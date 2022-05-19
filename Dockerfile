FROM golang:1.18.1-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download port-domain-service

COPY . .  

RUN mkdir -p /app

RUN go build -o /app/service ./cmd/app/main.go

ENTRYPOINT ["/app/service"]