FROM golang:1.23.1-alpine AS builder
WORKDIR /app
RUN apk add --no-cache \
    git \
    ca-certificates \
    tzdata \
    gcc \
    musl-dev
RUN adduser -D -s /bin/sh -u 1001 appuser
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main main.go
RUN ./main --version 2>/dev/null || echo "Binary built successfully"

FROM gcr.io/distroless/static-debian12:nonroot
RUN apk --no-cache add \
    ca-certificates \
    curl \
    tzdata
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /app/main
USER nonroot:nonroot
WORKDIR /app
EXPOSE 8080
LABEL maintainer="todo-app-team"
LABEL version="1.0"
LABEL description="Todo Application with Gin and GORM"
ENTRYPOINT ["/app/main"]