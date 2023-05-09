#!/bin/sh

docker build -t antlr -f Dockerfile.antlr .
docker run --rm -u "$(id -u):$(id -g)" -v "$(pwd):/antlr" antlr java -jar /antlr.jar SuperMakeFile.g4
