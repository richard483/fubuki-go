FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./fbk-go ./main.go


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/fbk-go .
EXPOSE 8080
ENTRYPOINT ["./fbk-go"]