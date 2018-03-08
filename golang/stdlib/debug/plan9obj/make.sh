#!/bin/bash

GOOS=plan9 GOARCH=amd64 go build -o hello &&

go build &&
./plan9obj
