#!/bin/bash
HOST=${HOST}
# register an info
curl -s ${HOST}/register -H 'Content-Type: application/json' -d '{"org":"https://www.hatena.com","user":"kawanos","redirectPath":"foo","comment":"test"}'
# infos for user
curl -s ${HOST}/info/user -H 'Content-Type: application/json'
# create user
curl -s ${HOST}/user -H 'Content-Type: application/json' -d '{"user":"tako","notify_to":"slack#log"}'
# delete user
curl -s ${HOST}/user -H 'Content-Type: application/json' -d '{"user":"tako"}' -X DELETE
# get path
curl -s ${HOST}/get/foo -H 'Content-Type: application/json'