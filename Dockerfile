FROM golang:1.24-alpine AS builder

# Uncomment the following line if you need access to the timezone data.
RUN apk add --no-cache tzdata ca-certificates

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Run tests and generate coverage report.
# Exit with non-zero status code if tests fail.
RUN go test ./... -coverprofile=coverage.out || exit 1

# Use -tags timetzdata to include the timezone data in the binary.
RUN go build -ldflags="-s -w" -tags timetzdata -o main cmd/api/main.go

FROM alpine:3.21.3


# Copy CA certificates for SSL verification
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder /build/main /
COPY pkg/config/db.migrations ./pkg/config/db.migrations/

# Command to run when starting the container.
ENTRYPOINT ["/main"]