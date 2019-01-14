#!/bin/bash

cd web
rm -rf dist; rm -rf node_modules; rm -f yarn.lock
yarn install && ng build --prod

cd dist/gorum && go-bindata -o ../../../internal/pkg/webassets/webassets.go -pkg webassets .
cd ../../../assets && go-bindata -o ../internal/pkg/assets/assets.go -pkg assets .

cd .. && rm -f gorum; rm -f gorum.asc;
GOOS=linux GOARCH=amd64 go build && gpg2 -a --detach-sign gorum
