#!/bin/bash
#check OS
platform='unknown'
unamestr=`uname`
isDevContainer="${DEVENV}"

if [[ "$unamestr" == 'Linux' ]]; then
    platform='linux'
elif [[ "$unamestr" == 'Darwin' ]]; then
    platform='macOS'
else
    echo "Platform not supported"
    exit 1
fi

version=$(go version)
regex='([1-9]|[1-9]{1}[0-9]{1}).([1-9]|[1-9]{1}[0-9]{1}).([1-9]|[1-9]{1}[0-9]{1})'
reqNM=0
if [[ $version =~ $regex ]]; then 
        if [ "${BASH_REMATCH[1]}" -lt 1 ]; then
            reqNM=1
        fi
        if [ "${BASH_REMATCH[2]}" -lt 12 ]; then
            reqNM=1
        fi
fi

if [ $reqNM -eq 1 ]; then
    echo 'Requirement not met!'
    echo 'Require Go version greater than 1.12'
    exit 0
fi

echo "Doing some clean-up"
# get vendor
rm -rf ./vendor
dep ensure -v &> ./get_dep.log


# Linux platform
if [[ "$platform" == 'linux' ]]; then

if [[ "$isDevContainer" == 'true' ]]; then

ln -s /v8/include/ ./vendor/github.com/augustoroman/v8/include &> ./get_dep.log
ln -s /v8/lib/ ./vendor/github.com/augustoroman/v8/libv8 &> ./get_dep.log

else

rm -rf ./deplib
mkdir -p ./deplib/v8/c/linux/data
(
    set -e
    curl https://rubygems.org/downloads/libv8-6.3.292.48.1-x86_64-linux.gem > ./deplib/v8/c/libv8.gem
    tar -xf ./deplib/v8/c/libv8.gem -C ./deplib/v8/c/linux/
    tar -xzf ./deplib/v8/c/linux/data.tar.gz  -C ./deplib/v8/c/linux/data
)
errorCode=$?
if [ $errorCode -ne 0 ]; then
    echo "Error getting v8(js) engine!"
    exit $errorCode
fi

ln -s -v $(pwd)/deplib/v8/c/linux/data/vendor/v8/include ./vendor/github.com/augustoroman/v8/include
ln -s -v $(pwd)/deplib/v8/c/linux/data/vendor/v8/out/x64.release ./vendor/github.com/augustoroman/v8/libv8

fi

fi

# macOS platform
if [[ "$platform" == 'macOS' ]]; then

rm -rf ./deplib
mkdir -p ./deplib/v8/c/macos/data
(
    set -e
    curl https://rubygems.org/downloads/libv8-6.3.292.48.1-universal-darwin-18.gem > ./deplib/v8/c/libv8.gem
    tar -xf ./deplib/v8/c/libv8.gem -C ./deplib/v8/c/macos/
    tar -xzf ./deplib/v8/c/macos/data.tar.gz  -C ./deplib/v8/c/macos/data
)
errorCode=$?
if [ $errorCode -ne 0 ]; then
    echo "Error getting v8(js) engine!"
    exit $errorCode
fi

ln -s -v $(pwd)/deplib/v8/c/macos/data/vendor/v8/include ./vendor/github.com/augustoroman/v8/include
ln -s -v $(pwd)/deplib/v8/c/macos/data/vendor/v8/out/x64.release ./vendor/github.com/augustoroman/v8/libv8

fi

echo "Get all dependencies done! Check get_dep.log for 'go get' error"

if [[ "$isDevContainer" == 'true' ]]; then
    tail -f /dev/null
fi