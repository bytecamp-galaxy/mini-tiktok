#!/bin/zsh
codegen -output ./errno_generated.go -type int base.go
codegen -doc -output ./errno_generated.md -type int base.go
