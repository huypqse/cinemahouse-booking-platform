# Build stage
FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app ./main.go

# Final stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
# Copy any static resources if needed, e.g.:
# COPY resource/public /app/resource/public
EXPOSE 8000
CMD ["./app"]