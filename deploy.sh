#!/bin/bash

IMAGE_NAME="hello-hertz"
IMAGE_TAG="latest"
CONTAINER_NAME="hello-hertz-app"
PORT=8888

echo "Building Docker image..."
docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .

echo "Stopping existing container (if exists)..."
docker stop ${CONTAINER_NAME} 2>/dev/null || true
docker rm ${CONTAINER_NAME} 2>/dev/null || true

echo "Starting new container..."
docker run -d \
  --name ${CONTAINER_NAME} \
  -p ${PORT}:${PORT} \
  --restart=always \
  ${IMAGE_NAME}:${IMAGE_TAG}

echo "Container started successfully!"
echo "Service is running at http://localhost:${PORT}"
docker logs ${CONTAINER_NAME}
