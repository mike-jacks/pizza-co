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
COPY ../../servers/order_management_server/ .

# Copy the clients
COPY ../../clients ./clients

COPY ../../servers/config ./servers/config

RUN ls -la /app

# Build the server application
RUN go build -o order_management_service .

# Expose the server port
EXPOSE 9000

# Command to run the server
CMD ["./order_management_service"]
