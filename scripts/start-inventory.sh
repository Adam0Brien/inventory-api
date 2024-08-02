#!/bin/bash
set -e
# Function to check if a command is available
source ./scripts/check_docker_podman.sh
${DOCKER} compose --profile relations-api --env-file ./scripts/.env -f ./docker-compose.yaml up -d