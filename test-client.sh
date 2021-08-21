#!/bin/bash
HOST=localhost:8081
# register an info
curl -s localhost:8081/register -H 'Content-Type: application/json' -d '{"org":"https://www.hatena.com","user":"kawanos","redirectPath":"foo","comment":"test"}'
# infos for user
curl -s localhost:8081/info/user -H 'Content-Type: application/json'
# create user
curl -s localhost:8081/user -H 'Content-Type: application/json' -d '{"user":"tako","notify":"slack#log"}'
# get path
curl -s localhost:8081/get/foo -H 'Content-Type: application/json'