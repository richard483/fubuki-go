FROM golang:1.24.1-alpine AS builder

# default 8080

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /out/main


FROM alpine:latest AS runner

WORKDIR /root/
COPY --from=builder /out/main .
CMD ["./main"]