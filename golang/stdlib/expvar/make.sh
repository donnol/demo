#!/bin/bash

pkill main &&
go build -o main &&
nohup ./main &
echo '测试 curl http://127.0.0.1:8822/'
