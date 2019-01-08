#!/bin/bash

# npm
cd web
ng update --all
npm install
cd ..

# dependencies
go get -u ./...
go mod vendor