# Use the official Golang image as the base image
FROM golang:1.17-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Set the working directory for the examples
WORKDIR /app/examples

# Copy the config.json file into the container
COPY examples/config.json .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/fishery .

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/fishery .

# Copy the config.json file into the container
COPY --from=builder /app/examples/config.json .

# Set the entry point for the container
ENTRYPOINT ["./fishery"]
