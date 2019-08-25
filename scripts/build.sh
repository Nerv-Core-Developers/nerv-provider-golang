#!/bin/bash
#check OS
unamestr=`uname`
if [[ "$unamestr" == 'Linux' ]]; then
    CGO_CXXFLAGS="-fno-rtti" go build -ldflags '-extldflags "-static -pthread -v"' -a -v -o ./build/server ./server/*.go
elif [[ "$unamestr" == 'Darwin' ]]; then
    go build -ldflags '-extldflags "-pthread -v"' -a -v -o ./build/server ./server/*.go 
else
    echo "Platform not supported"
    exit 1
fi

