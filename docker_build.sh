#!/bin/bash

# Set default tag
TAG=${TAG:-latest}

# Parse flags
while getopts ":t:" opt; do
  case $opt in
    t) TAG="v$OPTARG";;
    \?) echo "Invalid option: -$OPTARG"; exit 1;;
  esac
done

echo "About to build and output service files with version number: $TAG"
sleep 5

# Build the image
if ! docker build -t mikejacks/pizza-co-servers-inventory-service:$TAG -f ./servers/inventory_server/Dockerfile . --no-cache ; then
    echo "inventory-service build failed!"
    exit 1
fi

if ! docker build -t mikejacks/pizza-co-servers-order-management-service:$TAG -f ./servers/order_management_server/Dockerfile . --no-cache ; then
    echo "order-management-service build failed!"
    exit 1
fi

if ! docker tag mikejacks/pizza-co-servers-inventory-service:$TAG mikejacks/pizza-co-servers-inventory-service:latest ; then
    echo "inventory-service tag failed!"
    exit 1
fi

if ! docker tag mikejacks/pizza-co-servers-order-management-service:$TAG mikejacks/pizza-co-servers-order-management-service:latest ; then
    echo "order-management-service tag failed!"
    exit 1
fi

if ! docker push mikejacks/pizza-co-servers-order-management-service:$TAG ; then
    echo "Unable to push order-management-service:$TAG"
    exit 1
fi

if ! docker push mikejacks/pizza-co-servers-order-management-service:latest ; then
    echo "Unable to push order-management-service:latest"
    exit 1
fi

if ! docker push mikejacks/pizza-co-servers-inventory-service:$TAG ; then
    echo "Unable to push inventory-service:$TAG"
    exit 1
fi

if ! docker push mikejacks/pizza-co-servers-inventory-service:latest ; then
    echo "Unable to push inventory-service:latest"
    exit 1
fi

