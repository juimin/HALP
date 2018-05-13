#! /usr/bin/env bash

SCRIPT='
export TLSCERT=/etc/letsencrypt/live/staging.halp.derekwang.net/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/staging.halp.derekwang.net/privkey.pem
docker rm -f gateway
docker rmi -f d95wang/halp-gateway
docker pull d95wang/halp-gateway:latest
docker run -d \
-p 443:443 \
--name gateway \
--network appnet \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e REDISADDR=redisServer:6379 \
-e DBADDR=mongoServer:27017 \
-e SESSIONKEY=spUPraqUgethu4AF?x \
d95wang/halp-gateway
'
ssh root@159.65.250.147 "${SCRIPT}"