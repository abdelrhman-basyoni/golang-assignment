FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from your host machine to the container
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose a port
EXPOSE 3000

CMD ["./myapp"]
