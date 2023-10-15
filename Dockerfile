# Use the official Golang image as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from your host machine to the container
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose a port (if your Go application listens on a port)
EXPOSE 3000

# Command to run the Go application
CMD ["./myapp"]
