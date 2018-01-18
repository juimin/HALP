#! /usr/bin/env bash

export ADDR=localhost:4000
export SESSIONKEY=potato
export REDISADDR=localhost:6379
export DBADDR=localhost:27017
export TLSKEY=$(pwd)/tls/privkey.pem
export TLSCERT=$(pwd)/tls/fullchain.pem
go install && gateway