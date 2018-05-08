#!/bin/bash

# curl -v -X CONNECT -k -x https://localhost:8820 https://www.google.co.uk
# curl -v -X CONNECT -x localhost:8820 www.google.co.uk
curl -Lv --proxy https://localhost:8820 https://google.com # --proxy-cacert server.pem 