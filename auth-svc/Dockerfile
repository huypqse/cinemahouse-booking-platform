# Build stage
FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o auth-svc ./main.go

# Final stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/auth-svc .
# If you need static/config files, add lines like:
# COPY resource /app/resource
EXPOSE 8000
CMD ["./auth-svc"]