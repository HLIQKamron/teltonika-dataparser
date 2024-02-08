# Use the official Golang image as a base
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the entire contents of the current directory into the container's working directory
COPY . .

# Build the Go application

# Expose the port the application runs on
EXPOSE 8081

# Command to run the executable
CMD ["go","run","cmd/main.go"]
