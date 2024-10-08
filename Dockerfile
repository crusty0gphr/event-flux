FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o eventflux ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/eventflux .

EXPOSE 8080

# Command to run the executable
CMD ["./eventflux"]
