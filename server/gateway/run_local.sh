#! /usr/bin/env bash

# Run the docker images for redis (session storage)
# and the mongo db

# Check if the local redis is running
if [[ $(docker inspect -f {{.State.Running}} redislocal) -eq true ]]
then
   # Reset the mongo
   docker rm -f redislocal
fi

# Check if the local mongo is running
if [[ $(docker inspect -f {{.State.Running}} mongolocal) -eq true ]]
then
   # Reset the mongo
   docker rm -f mongolocal
fi

# Start new instance of the redis and mongo services
docker run -d -p 6379:6379 --name redislocal redis
echo "Redis Local started successfully"
docker run -d -p 27017:27017 --name mongolocal mongo
echo "Mongo Local started successfully"

# Set environment variables
export ADDR=localhost:4000
export SESSIONKEY=potato
export REDISADDR=localhost:6379
export DBADDR=localhost:27017
export TLSKEY=$(pwd)/tls/privkey.pem
export TLSCERT=$(pwd)/tls/fullchain.pem

# Run the server locally
go install && gateway