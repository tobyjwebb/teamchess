#!/bin/bash

echo "CURRENT_TIME $(date +%s)"
echo "RANDOM_HASH $(cat /dev/urandom | head -c16 | md5sum 2>/dev/null | cut -f1 -d' ')"
echo "STABLE_GIT_COMMIT $(git rev-parse HEAD)"
echo "DOCKER_REGISTRY $DOCKER_REGISTRY"
echo "DOCKER_REPOSITORY $DOCKER_REPOSITORY"

# Handle the '/' characters in branch names (and any non-alphanumeric characters):
echo "DOCKER_IMAGE_TAG $(echo -n "$DOCKER_IMAGE_TAG" | tr -c [:alnum:] _)"
