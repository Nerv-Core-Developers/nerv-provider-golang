#!/bin/sh
set -e
docker build -f docker/dev-image/dev.Dockerfile . -t nerv-provider-golang-devenv