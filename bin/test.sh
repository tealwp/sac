#!/bin/bash

# make coverage directory
mkdir -p ./coverage

# run testts for project with outputs going into the ./coverage directory
go test -coverprofile=./coverage/coverage.out ./...

# generate the html report
go tool cover -html=./coverage/coverage.out -o ./coverage/index.html
