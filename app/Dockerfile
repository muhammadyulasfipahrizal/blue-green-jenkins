FROM golang:1.22.2-alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:3.22

RUN adduser -D -H appuser

WORKDIR /app

COPY --from=builder /src/app .

USER appuser

EXPOSE 8080

ENTRYPOINT ["./app"]