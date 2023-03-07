# Use the official Golang image as a base
FROM golang:alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the application
RUN go build .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["go run main.go", "go run main.go Calculate"]

 