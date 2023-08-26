#!/usr/bin/env bash

docker build -t simple .
docker run -p 8080:8080 simple