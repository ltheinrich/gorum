#!/bin/bash

URL=$1
DATA=$2

curl -k -i -H 'Content-Type: application/json' -X POST -d "$DATA" https://localhost:1813/$URL
echo