#!/bin/bash

# build web-app
cd web
npm install
ng build --prod
cd ..

# assets binary data
cd assets
go-bindata -o ../internal/pkg/assets/assets.go -pkg assets .
cd ..

# web assets binary data
cd web/dist/gorum
go-bindata -o ../../../internal/pkg/webassets/webassets.go -pkg webassets .
cd ../../..

# linux amd64
GOOS=linux GOARCH=amd64 go build -o bin/gorum-linux-amd64 cmd/gorum/gorum.go

# linux arm64
GOOS=linux GOARCH=arm64 go build -o bin/gorum-linux-arm64 cmd/gorum/gorum.go

# linux armv6
GOOS=linux GOARCH=arm GOARM=6 go build -o bin/gorum-linux-armv6 cmd/gorum/gorum.go