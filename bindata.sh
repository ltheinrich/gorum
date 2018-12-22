#!/bin/bash

cd assets
go-bindata -o ../internal/pkg/assets/assets.go -pkg assets .
cd ..
