#!/bin/bash

# Define the image name and tag
IMAGE_NAME="docker-forum"
IMAGE_TAG="latest"

# Build the Docker image
docker build -t "${IMAGE_NAME}:${IMAGE_TAG}" .

# Run a Docker container
docker run -d -p 8080:8080 --name "${IMAGE_NAME}" "${IMAGE_NAME}:${IMAGE_TAG}"

# Display container information
docker ps --filter "name=${IMAGE_NAME}"

# Provide instructions on stopping the container
echo "To stop the container, run: docker stop ${IMAGE_NAME}"
