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
GOOS=linux GOARCH=amd64 go build -o bin/gorum-linux-amd64 main.go

# linux arm64
GOOS=linux GOARCH=arm64 go build -o bin/gorum-linux-arm64 main.go

# linux armv6
GOOS=linux GOARCH=arm GOARM=6 go build -o bin/gorum-linux-armv6 main.go

# freebsd amd64
GOOS=freebsd GOARCH=amd64 go build -o bin/gorum-freebsd-amd64 main.go

# freebsd armv6
GOOS=freebsd GOARCH=arm GOARM=6 go build -o bin/gorum-freebsd-armv6 main.go
