#!/bin/bash
HOST=${HOST}
# register an info
curl -s ${HOST}/register -H 'Content-Type: application/json' -d '{"org":"https://www.hatena.com","user":"tako","redirectPath":"foo","comment":"test"}'
# infos for user
curl -s ${HOST}/info/tako -H 'Content-Type: application/json'
# create user
curl -s ${HOST}/user/tako -H 'Content-Type: application/json' -d '{"notify_to":"slack#log"}' -X POST
# delete user
curl -s ${HOST}/user/tako -H 'Content-Type: application/json' -X DELETE
# get path
curl -s ${HOST}/get/foo -H 'Content-Type: application/json'