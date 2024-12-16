# Base image
FROM golang:1.23.3

# Set working directory
WORKDIR /app

# Copy Go modules and source code
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN go build -o main .

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./main"]
