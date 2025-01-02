# Start from the latest golang base image
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first; they are less frequently changed than source code, so Docker can cache this layer
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]




# # Start from a more recent official Go image
# FROM golang:1.22-alpine AS builder

# # Set the working directory inside the container
# WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# # Download all dependencies
# RUN go mod download

# # Copy the entire project
# COPY . .

# # Build the application
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# # Start a new stage from scratch
# FROM alpine:latest  

# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# # Copy the pre-built binary file from the previous stage
# COPY --from=builder /app/main .

# # Command to run the executable
# CMD ["./main"]