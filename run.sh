#!/bin/bash
#instead of test can be set real dockerhub repository name
docker run -ti --rm -v ${PWD}:/home/user test/build-cheatsheeter-image
