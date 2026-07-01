FROM golang:1.26.4-alpine AS builder

WORKDIR /app

COPY go.mod go.su[m] ./

RUN go mod download

COPY . .

RUN go build -o ascii-art-web .


FROM alpine:3.22

WORKDIR /app

LABEL org.opencontainer.image.title="ascii-art-web"
LABEL org.opencontainer.image.version="1.0"
LABEL org.opencontainer.image.source="github.com/HerPerception/my-go-playbook/ascii-web"

COPY --from=builder /app/ascii-art-web .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/banners ./banners

EXPOSE 8080

CMD ["./ascii-art-web"]