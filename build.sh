#!/bin/bash

# build web-app
cd web
rm -rf dist
rm -rf node_modules
rm -f package-lock.json
npm install
ng build --prod
cd ..

# web assets binary data
cd web/dist/gorum
go-bindata -o ../../../internal/pkg/webassets/webassets.go -pkg webassets .
cd ../../..

# assets binary data
cd assets
go-bindata -o ../internal/pkg/assets/assets.go -pkg assets .
cd ..

# linux amd64
GOOS=linux GOARCH=amd64 go build -o bin/gorum-linux-amd64

# linux arm64
GOOS=linux GOARCH=arm64 go build -o bin/gorum-linux-arm64

# linux armv6
GOOS=linux GOARCH=arm GOARM=6 go build -o bin/gorum-linux-armv6

# freebsd amd64
GOOS=freebsd GOARCH=amd64 go build -o bin/gorum-freebsd-amd64

# freebsd armv6
GOOS=freebsd GOARCH=arm GOARM=6 go build -o bin/gorum-freebsd-armv6