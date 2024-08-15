# Step 1: Build the Go binary
FROM golang:1.22.6-alpine

# Set the working directory
WORKDIR /app

# Copy the Go modules and dependencies
COPY go.mod go.sum ./

# Download all Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]