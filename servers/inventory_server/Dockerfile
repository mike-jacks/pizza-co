# Use an official Go image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY ../../../go.mod ../../../go.sum ./

RUN go mod download

# Copy the generated protobuf files
COPY ../../../generated ./generated

# Copy the server code
COPY ../../servers/inventory_server/ .

COPY ../../servers/config ./servers/config

# Build the server application
RUN go build -o inventory_service .

# Expose the server port
EXPOSE 9100

# Command to run the server
CMD ["./inventory_service"]
