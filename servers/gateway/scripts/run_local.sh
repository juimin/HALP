#! /usr/bin/env bash
./reset_local_services.sh

cd ../
# Set environment variables
export ADDR=localhost:4000
export SESSIONKEY=potato
export REDISADDR=localhost:6379
export DBADDR=localhost:27017
export TLSKEY=$(pwd)/tls/privkey.pem
export TLSCERT=$(pwd)/tls/fullchain.pem

# Run the server locally
go install && gateway