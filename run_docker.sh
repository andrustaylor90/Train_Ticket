#!/bin/bash

# Shell script located in the root of TrainTicketServices directory

# Define Docker variables
IMAGE_NAME="grpc-train-ticket-server"
CONTAINER_NAME="grpc-train-ticket-server-instance"
PORT=50051

# Stop and remove the old container
echo "Stopping existing container..."
docker stop $CONTAINER_NAME
docker rm $CONTAINER_NAME

# Build the Docker image
echo "Building Docker image..."
docker build -t $IMAGE_NAME .

# Run the Docker container
echo "Running Docker container..."
docker run -d --name $CONTAINER_NAME -p $PORT:50051 $IMAGE_NAME

echo "Container $CONTAINER_NAME is now running on port $PORT"
