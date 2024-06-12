# Use an existing Go image as the base image
FROM golang:1.22-alpine

# Set the working directory in the container to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the project files to the container
COPY . .

# Build the Go Gin project
RUN apk add --no-cache build-base librdkafka-dev pkgconf
RUN make build

# Set the command to run when the container starts
CMD ["./webserver"]

