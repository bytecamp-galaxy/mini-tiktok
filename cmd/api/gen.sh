#!/bin/zsh
hz new --module github.com/bytecamp-galaxy/mini-tiktok/api-server --idl ../../idl/api.thrift --service api --exclude_file main.go --exclude_file router.go && rm -f .gitignore go.mod