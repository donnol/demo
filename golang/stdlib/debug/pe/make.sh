#!/bin/bash

GOOS=windows GOARCH=amd64 go build -o hello

go build
./pe
