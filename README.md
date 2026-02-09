# golang-technical-assesment-v2

## Overview
This is a Golang-based microservice designed to report the energy consumption of a set of electricity meters.

## Prerequisites
- Go 1.24.x
- Docker

## Configuration

[.env.example](.env.example) contains example variables which you can use to start the service locally.

## Running the application

### Local development

To run the application locally, you can use the following commands:
```bash
docker compose up -d
go run cmd/api/main.go
```

This will start the application on `localhost:8181`.

You can check [the API specification file](./internal/interface/Api/spec/openapi.yaml) to know how to interact with the API.

## Testing

Run the tests using:
```sh
go test -v ./...
```

<!-- TODO: k6 integration tests -->
