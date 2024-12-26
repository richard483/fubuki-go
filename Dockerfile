FROM golang:1.22.2-alpine AS builder

# default 8080
ARG PORT=8080

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./fbk-go ./main.go


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/fbk-go .
EXPOSE $PORT
ENTRYPOINT ["./fbk-go"]