FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o my-app .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/my-app .
CMD ["./my-app"]
