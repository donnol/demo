#!/bin/bash

# 生成您的专用密钥和公用证书
openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out cert.pem

# 将密钥和证书组合在 PKCS#12 (P12) 捆绑软件中
openssl pkcs12 -inkey key.pem -in cert.pem -export -out cert.p12
