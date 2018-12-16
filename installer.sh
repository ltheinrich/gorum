#!/bin/bash

ARCH=$(arch)
ARCHITECTURE=""

if [ $ARCH == "x86_64" ]; then
    ARCHITECTURE="amd64"
elif [ $ARCH == "armv8*" ] || [ $ARCH == "arm64*" ]; then
    ARCHITECTURE="arm64"
elif [ $ARCH == "armv7*" ]; then
    ARCHITECTURE="armv7"
elif [ $ARCH == "arm*" ]; then
    ARCHITECTURE="armv6"
elif [ $ARCH == "i*86" ]; then
    ARCHITECTURE="386"
else
    echo "unsupported architecture"
    exit
fi

if [ "$1" != "" ]; then
        # executable
        wget -O gorum-linux https://github.com/ltheinrich/gorum/releases/download/v$1/gorum-linux-$ARCHITECTURE
        chmod +x gorum-linux

        # resources
        wget https://github.com/ltheinrich/gorum/releases/download/v$1/resources.tar.gz
        tar xfvz resources.tar.gz
else
    echo "use ./installer.sh VERSION(x.y.z)"
    exit
fi