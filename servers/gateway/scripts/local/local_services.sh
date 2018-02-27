#! /usr/bin/env bash

# Run the docker images for redis (session storage)
# and the mongo db

# Get to the main file one dir up to do the install and start gateway

# Check if the local redis is running
if [[ $(docker inspect -f {{.State.Running}} redisLocal) -eq true ]]
then
   # Reset the mongo
   docker rm -f redisLocal
fi

# Check if the local mongo is running
if [[ $(docker inspect -f {{.State.Running}} mongoLocal) -eq true ]]
then
   # Reset the mongo
   docker rm -f mongoLocal
fi

# Start new instance of the redis and mongo services
docker run -d -p 6379:6379 --name redisLocal redis
echo "Redis Local started successfully"
docker run -d -p 27017:27017 --name mongoLocal mongo
echo "Mongo Local started successfully"