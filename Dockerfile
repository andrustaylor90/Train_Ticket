# Dockerfile located in the root of TrainTicketServices directory

# Use the official Golang image to create a build artifact.
FROM golang:1.19 as builder

# Enable Go modules.
ENV GO111MODULE=on

WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./server/server.go

# Use a small Alpine Linux image to run the server.
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the server binary from the builder image.
COPY --from=builder /app/server .

# Expose server port (default gRPC port).
EXPOSE 50051

# Command to run the binary.
CMD ["./server"]
