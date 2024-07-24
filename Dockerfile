# Use an official Go image with Alpine as the base image
FROM golang:1.22.4-alpine

# Install required packages for building and running Go applications
RUN apk --no-cache add build-base ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# download dependencies
RUN go mod download

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

# Expose the port on which the application will run
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
