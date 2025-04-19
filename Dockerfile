FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./fbk-go ./main.go


FROM alpine:latest AS runner

RUN apk add --no-cache bash

WORKDIR /app
COPY --from=builder /app/fbk-go .
EXPOSE $PORT
ENTRYPOINT ["./fbk-go"]