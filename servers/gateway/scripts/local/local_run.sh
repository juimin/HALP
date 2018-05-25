#! /usr/bin/env bash
./local_services.sh

export GATEWAYPATH=$GOPATH/src/github.com/JuiMin/HALP/servers/gateway
cd $GATEWAYPATH
# Set environment variables
export ADDR=localhost:4000
export SESSIONKEY=potato
export REDISADDR=localhost:6379
export DBADDR=localhost:27017
export TLSCERT=$GATEWAYPATH/tls/fullchain.pem
export TLSKEY=$GATEWAYPATH/tls/privkey.pem

# Run the server locally
go install && gateway
