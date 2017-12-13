#!/bin/sh
curl -fsSL get.docker.com -o get-docker.sh

# 要翻墙 https_proxy=https://127.0.0.1:1080
sudo sh get-docker.sh
