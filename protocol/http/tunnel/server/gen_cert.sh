#!/bin/bash

# 生成证书，后面需要添加 --host www.yoursite.com
go run /usr/local/go/src/crypto/tls/generate_cert.go $@