FROM golang:1.23.0-alpine3.20 as builder
WORKDIR /app

# Install gRPC health probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.13 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

# Install application dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the application
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./main.go

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe
COPY --from=builder /main ./main