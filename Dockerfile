# Use the official Go image with Alpine as the base
FROM golang:1.20-alpine

# Install bash
RUN apk add --no-cache curl bash git

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Specify the command to run the application
CMD ["./main"]
