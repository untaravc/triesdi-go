# Use Go official image as a base
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy Go modules & download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app source code
COPY . .
COPY .env .env

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["/app/main"]
