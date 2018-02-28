#! /usr/bin/env bash

# Move to the main folder
cd ../

docker rm -f gateway
docker rmi -f d95wang/halp-gateway

set -e
echo "Building Halp Gateway..."
# This mighewt need to change based on what Operating System you use
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