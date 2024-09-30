#!/bin/bash

# run testts for project with outputs going into the ./coverage folder
go test -coverprofile=./coverage/coverage.out ./...

# generate the html report
go tool cover -html=./coverage/coverage.out -o ./coverage/index.html
