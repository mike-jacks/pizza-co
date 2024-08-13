#!/bin/bash

# Define variables
POD_NAME="svc/order-management-service"
LOCAL_PORT=9000
REMOTE_PORT=9000

# Function to start port forwarding
start_port_forward() {
  echo "Starting port forwarding..."
  kubectl port-forward $POD_NAME $LOCAL_PORT:$REMOTE_PORT &
  PORT_FORWARD_PID=$!
}

# Initial port forwarding
start_port_forward

# Loop to monitor and restart if it fails
while true; do
  # Check if the port forwarding process is running
  if ! ps -p $PORT_FORWARD_PID > /dev/null; then
    echo "Port forwarding process has stopped. Restarting..."
    start_port_forward
  fi

  # Sleep for a short duration before checking again
  sleep 1
done