# Stage 1: Build GitVerse binary
FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/gitverse ./cmd/gitverse

# Stage 2: Minimal scratch runtime
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/gitverse /bin/gitverse
ENTRYPOINT ["/bin/gitverse"]
