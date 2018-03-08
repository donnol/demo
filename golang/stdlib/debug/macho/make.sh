#!/bin/bash

GOOS=darwin GOARCH=amd64 go build -o hello

go build
./macho
