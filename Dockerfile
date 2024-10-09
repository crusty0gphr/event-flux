FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o eventflux ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/eventflux .
COPY ./scripts/wait-for-it.sh /scripts/wait-for-it.sh
RUN apk add --no-cache bash
RUN chmod +x /scripts/wait-for-it.sh
EXPOSE 8080
CMD ["./eventflux"]