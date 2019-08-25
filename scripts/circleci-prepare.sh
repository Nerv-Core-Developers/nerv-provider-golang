#!/bin/bash
dep ensure -v
ln -s /v8/include/ ./vendor/github.com/augustoroman/v8/include
ln -s /v8/lib/ ./vendor/github.com/augustoroman/v8/libv8

# CGO_CXXFLAGS="-fno-rtti" go build -ldflags '-extldflags -static' -a -v -o server ./server/*.go