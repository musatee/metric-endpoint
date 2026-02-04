# Metric Endpoint

A simple HTTP server that exposes an API endpoint which returns error rate metrics.

## Features

- Simple HTTP server using Go's `net/http` package
- RESTful API endpoint at `/apis/error`
- Returns JSON response with error rate data
- Request logging with method, path, and response time
- Dockerized deployment with multi-stage build

## API Endpoint

### GET `/apis/error`

Returns error rate metrics in JSON format.

**Response:**
```json
{
  "error_rate": 20.5,
  "message": "Current error rate"
}
```

## Running Locally

### Prerequisites
- Go 1.25.6 or higher

### Run the server
```bash
go run main.go
```

The server will start on port `8080`.

### Test the endpoint
```bash
curl http://localhost:8080/apis/error
```

## Docker Deployment

### Build the Docker image
```bash
docker build -t metric-endpoint .
```

### Run the container
```bash
docker run -p 8080:8080 metric-endpoint
```