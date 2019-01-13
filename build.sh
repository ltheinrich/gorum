#!/bin/bash

cd web
rm -rf dist; rm -rf node_modules; rm -f package-lock.json
npm install && ng build --prod

cd dist/gorum && go-bindata -o ../../../internal/pkg/webassets/webassets.go -pkg webassets .
cd ../../../assets && go-bindata -o ../internal/pkg/assets/assets.go -pkg assets .

cd .. && rm -rf bin
GOOS=linux GOARCH=amd64 go build -o bin/gorum-linux-amd64
GOOS=freebsd GOARCH=amd64 go build -o bin/gorum-freebsd-amd64
cd bin && gpg2 -a --detach-sign gorum-linux-amd64 && gpg2 -a --detach-sign gorum-freebsd-amd64
