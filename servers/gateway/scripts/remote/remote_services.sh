#! /usr/bin/env bash

# Set up the docker network and the docker apps
SCRIPT='
docker network create appnet

if [[ $(docker inspect -f {{.State.Running}} redisServer) -eq true ]]
then
   docker rm -f redisServer
fi

if [[ $(docker inspect -f {{.State.Running}} mongoServer) -eq true ]]
then
   docker rm -f mongoServer
fi

docker run -d -p 6379:6379 --name redisServer --network appnet redis
docker run -d -p 27017:27017 --name mongoServer --network appnet mongo
'
ssh root@159.65.250.147 "${SCRIPT}"