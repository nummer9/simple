#!/usr/bin/env bash

echo "---------------------| Starting JENKINS-BUILD Nr. ${BUILD_NUMBER} |------------------------"

export GOPATH="$JENKINS_HOME/workspace/go"
go get -u github.com/stretchr/testify
go test simple
go build simple
