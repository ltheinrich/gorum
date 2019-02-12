#!/bin/bash

cd assets && go-bindata -debug -o ../internal/pkg/assets/assets.go -pkg assets . && cd ..
clear && go run main.go
