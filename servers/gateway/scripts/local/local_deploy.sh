#! /usr/bin/env bash
export GATEWAYPATH=$GOPATH/src/github.com/JuiMin/HALP/servers/gateway/
   
if [[ $(docker inspect -f {{.State.Running}} gateway) -eq true ]]
then
   # Reset the mongo
   docker rm -f gateway
fi

docker rmi -f d95wang/halp-gateway

docker pull d95wang/halp-gateway:latest

docker run -d \
-p 443:443 \
--network appnet \
--name gateway \
-v $GATEWAYPATH/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
-e REDISADDR=redisLocal:6379 \
-e DBADDR=mongoLocal:27017 \
-e SESSIONKEY='spUPraqUgethu4AF?x' \
d95wang/halp-gateway
