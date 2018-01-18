#! /usr/bin/env bash

# Move to the main folder
cd ../

set -e
echo "Gateway Building..."
# This might need to change based on what Operating System you use
CGO_ENABLED=0 go build -a

# You might want to change this later
# to your own docker hub in case yow want to build at the same time

# Build gateway
docker build -t d95wang/halp-gateway .

# Push the gateway to dockerhub
docker push d95wang/halp-gateway

# Clean something?
go clean
echo "Gateway Built"