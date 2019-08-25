#!/bin/sh
set -e
docker build -f docker/production/prod.Dockerfile -t nerv-provider-golang