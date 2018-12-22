#!/bin/bash

# dependencies
go get -u ./...
go mod vendor

# binary data
cd assets
go-bindata -o ../internal/pkg/assets/assets.go -pkg assets .
cd ..