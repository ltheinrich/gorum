#!/bin/bash

METHOD=$1
URL=$2
DATA=$3

if [ $1 != "POST" ] && [ $1 != "GET" ] && [ $1 != "PUT" ] && [ $1 != "DELETE" ]
    then
        METHOD="POST"
        URL=$1
        DATA=$2
fi

curl -k -i -H 'Content-Type: application/json' -X $METHOD -d "$DATA" https://localhost:1813/$URL
echo