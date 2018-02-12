#!/bin/sh
curl -fsSL get.docker.com -o get-docker.sh

# 直接获取失败的话，使用代理试试 https_proxy=https://127.0.0.1:1080
sudo sh get-docker.sh
