#! /usr/bin/env bash
cd $CAPSTONE/servers/gateway
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
rm coverage.out coverage.html
