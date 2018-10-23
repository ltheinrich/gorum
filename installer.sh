#!/bin/bash

if [ "$1" != "" ] && [ "$2" != "" ]; then
        # executable
        wget -O gorum-linux https://github.com/lheinrichde/gorum/releases/download/v$1/gorum-linux-$2
        chmod +x gorum-linux

        # resources
        wget https://github.com/lheinrichde/gorum/releases/download/v$1/resources.tar.gz
        tar xfvz resources.tar.gz
else
    echo "use ./installer.sh VERSION(x.y.z) ARCHITECTURE(amd64,arm64,armv6,armv7,386)"
fi