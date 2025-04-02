#!/bin/bash

# remove previous build output
rm -rf bin

# remove build cache
go clean -cache

# module depndencies
go mod tidy -C src

# build executable binary
go build -C src  -o ../bin/localfs  -ldflags="-s -w" -x
