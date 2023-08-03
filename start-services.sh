#!/bin/bash

# Function to start a service
start_service() {
  service=$1
  echo "Starting $service..."
  cd $service/cmd
  go run main.go &
  cd ../..
}

# Start each service
start_service auth-service
start_service user-management-service
start_service video-streaming-service
start_service video-upload-service

echo "All services started"
