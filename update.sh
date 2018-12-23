#!/bin/bash

# npm
cd web
ng update --all
npm install
cd ..

# dependencies
go get -u ./...
go mod vendor

# assets binary data
cd assets
go-bindata -o ../internal/pkg/assets/assets.go -pkg assets .
cd ..

# web assets binary data
cd web/dist/gorum
go-bindata -o ../../../internal/pkg/webassets/webassets.go -pkg webassets .
cd ../../..