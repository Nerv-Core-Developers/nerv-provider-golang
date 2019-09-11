#!/bin/sh
set -e
#check OS
platform='unknown'
unamestr=`uname`
if [[ "$unamestr" == 'Linux' ]]; then
    platform='linux'
elif [[ "$unamestr" == 'Darwin' ]]; then
    platform='macOS'
else
    platform='windows'
fi

echo "[Linux/macOS dev-env deploy script v1.0 (can be used for Docker Toolbox on Windows)]"
#check whether docker available
command -v docker >/dev/null 2>&1 || { echo >&2 "Docker is not available! Please install docker  Aborting."; exit 1; }

if [[ "$platform" != 'windows' ]]; then

#Deploy stage
printf "Downloading docker images (about 400MB) ... "
(
    set -e
    docker pull alpine:3.7 >/dev/null & \
    docker pull naokichau/nerv-dev:latest >/dev/null
)
errorCode=$?
if [ $errorCode -ne 0 ]; then
    echo "\033[31m\nDeploy failed ✗\033[0m"
    exit $errorCode
fi
echo "\033[32mdownloaded ✔\033[0m"
printf "Deploying dev environment ... "
(
    set -e
    COMPOSE_PROJECT_NAME=nervdev
    docker-compose up -d &>./deploy_logs.log
)
errorCode=$?
if [ $errorCode -ne 0 ]; then
    echo "\033[31m\nDeploy failed ✗\033[0m"
    exit $errorCode
fi
echo "\033[32mdeployed ✔\033[0m"

echo "Dev environment is ready! 🎉  🎉  🎉"
echo "To start a node: run '\033[34mdocker exec -it nerv-provider-1 sh\033[0m' and run '\033[34m./scripts/build.sh && ./build/server -config /envs/env1.yaml\033[0m'."
echo "Open 3 more terminal and run '\033[34mdocker exec -it nerv-provider-<NODE_NUMBER> sh\033[0m' -> '\033[34m./server -config /envs/env2.yaml\033[0m'."
echo "\033[31m\nNOTE:\033[0m   - You might need to wait for dev container to get all the Go dependencies before you can 'Go for a Run'.
        - If you install new go package you need to install it inside the container."
echo "\nHappy coding 🤘"
exit 0

else
#####WINDOWS SCRIPT######
#Deploy stage
printf "Windows dev-env is currently not available!"

# printf "Downloading docker images (about 380MB) ... "
# (
#     set -e
#     docker pull alpine:3.8 >/dev/null & \
#     docker pull golang:1.11rc1-alpine3.8 >/dev/null
# )
# errorCode=$?
# if [ $errorCode -ne 0 ]; then
#     echo -e "\e[31m\nDeploy failed\e[0m"
#     exit $errorCode
# fi
# echo -e "\e[32mdownloaded\e[0m"
# printf "Deploying dev environment ... "
# (
#     set -e
#     docker-compose up -d &>./deploy_logs.log
# )
# errorCode=$?
# if [ $errorCode -ne 0 ]; then
#     echo -e "\e[31m\nDeploy failed\e[0m"
#     exit $errorCode
# fi
# echo -e "\e[32mdeployed\e[0m"

# echo -e "Dev environment is ready!"
# echo "To start a node: run '\e[34mdocker exec -it nerv-node-NODE_NUMBER sh\e[0m' and run '\e[34mcd /go/src/github.com/Nerv-Core-Developers/nerv-provider-golang\e[0m' then run '\e[34mgo run server.go\e[0m'."
# echo "\e[31m\nNOTE:\e[0m   - You might need to wait for dev containers to get all the Go dependencies before you can 'Go for a Run'.
#         - If you install new go package you need to install it inside the container."
# echo -e "\nSadly coding :_("
# exit 0


fi