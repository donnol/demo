#!/bin/bash

# 生成性能报告
go test -v -bench=. -cpuprofile=profile.out

# 查看报告
# go tool pprof -http=:8080 profile.out
pprof -http=:8080 profile.out
