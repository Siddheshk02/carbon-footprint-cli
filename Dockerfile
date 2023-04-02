# Use the official Golang image as the base image
FROM golang:alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the CLI binary
RUN go build -o carbon-footprint-cli main.go

# Set the default command for the container to run the main command
CMD ["./carbon-footprint-cli"]

# Define the second command as an entrypoint
ENTRYPOINT ["./carbon-footprint-cli", "Calculate"]
