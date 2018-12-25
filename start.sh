#!/bin/bash

# assets binary data
cd assets
go-bindata -o ../internal/pkg/assets/assets.go -pkg assets .
cd ..

# start
clear
go run main.go