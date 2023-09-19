# Use the official Golang image as a base image
FROM golang:1.21

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download Go module dependencies
RUN go mod download

# Copy the rest of your application source code to the working directory
COPY . .

# Install PostgreSQL and its client libraries
RUN apt-get update && apt-get install -y postgresql postgresql-client

# Build your Go application
RUN go build -o ./out/sever-api

# Expose port 8080 (if your application listens on this port)
EXPOSE 8080

# Set the command to run your application when the container starts
CMD ["./out/sever-api"]
