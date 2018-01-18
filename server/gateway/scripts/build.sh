#! /usr/bin/env bash
set -e
echo "Gateway Building..."
# This might need to change based on what Operating System you use
CGO_ENABLED=0 go build -a

# You might want to change this later
# to your own docker hub in case yow want to build at the same time

# Build gateway
docker build -t d95wang/gateway .

# Push the gateway to dockerhub
docker push d95wang/gateway

# Clean something?
go clean
echo "Gateway Built"