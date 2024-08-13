# Use an official Go image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY ./go.mod ./go.sum ./

RUN go mod download

# Copy the generated protobuf files
COPY ./common/ports/grpc/v1 ./common/ports/grpc/v1
COPY ./inventory_service/ports/grpc/v1 ./inventory_service/ports/grpc/v1
# Copy the server code
COPY ./order_management_service/. ./order_management_service/

# Copy the clients
COPY ./clients ./clients

COPY ./config ./config

# Build the server application
RUN go build -o order_management_service ./order_management_service/.

# Expose the server port
EXPOSE 9000

# Command to run the server
CMD ["./order_management_service"]
