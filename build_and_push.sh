#!/bin/bash

# Define your Docker Hub username
DOCKER_HUB_USERNAME="nurulhadii"

# List of services and their directories
services=(
  "spk-services"
  "sph-services"
  "bank-services"
  "user-management"
  "project-management"
  "breakdown-services"
)

# Build and push each service
for service in "${services[@]}"; do
  echo "Building and pushing $service..."
  docker build -t $DOCKER_HUB_USERNAME/$service:latest ./services/$service
  docker push $DOCKER_HUB_USERNAME/$service:latest
done

# Build and push the API Gateway
echo "Building and pushing api-gateway..."
docker build -t $DOCKER_HUB_USERNAME/api-gateway:latest ./api-gateway
docker push $DOCKER_HUB_USERNAME/api-gateway:latest

echo "All services have been built and pushed successfully!"
