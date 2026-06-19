FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ascii-art-web-dockerize .

FROM alpine:3.22

WORKDIR /app

LABEL org.opencontainers.image.title="ascii-art-web-dockerize"
LABEL org.opencontainers.image.version="1.0"
LABEL org.opencontainers.image.source="https://github.com/HerPerception/my-go-playbook"

COPY --from=builder /app/ascii-art-web-dockerize .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/standard.txt .
COPY --from=builder /app/shadow.txt .
COPY --from=builder /app/thinkertoy.txt .

EXPOSE 8080

CMD ["./ascii-art-web-dockerize"]
