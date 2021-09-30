.PHONY:build

install_gotip:
	go get golang.org/dl/gotip
	export https_proxy=http://127.0.0.1:51837
	CGO_ENABLED=0 gotip download dev.typeparams

build:
	# gotip build -gcflags=-G=3 main.go
	gotip build main.go

run:
	gotip run main.go
