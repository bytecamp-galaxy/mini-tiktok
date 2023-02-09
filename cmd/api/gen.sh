#!/bin/zsh
# set working directory to cmd/api
hz new --module github.com/bytecamp-galaxy/mini-tiktok/cmd/api --idl ../../idl/api/api_service.thrift --service api --exclude_file main.go --exclude_file router.go && rm -f .gitignore go.mod
swag init