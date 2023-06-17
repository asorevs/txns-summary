# Use the official Go Docker image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the entire repository into the container
COPY . .

# Build the Go application
RUN go build -o txn-summary

# The setup.sh script will handle environment variable setup and Docker run command
CMD ["./txn-summary"]
