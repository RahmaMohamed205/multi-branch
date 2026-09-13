# --- Stage 1: Build ---
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# --- Stage 2: Final image ---
FROM scratch

COPY --from=builder /app/server /server

EXPOSE 5000

CMD ["/server"]
