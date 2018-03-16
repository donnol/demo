package main

import (
	"go/build"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("%+v\n%+v\n", build.Default, build.Default.SrcDirs())

	arch, err := build.ArchChar(build.Default.GOARCH)
	if err != nil {
		// log.Fatal(err)
	}
	log.Printf("%s\n", arch)

	pkg, err := build.Import("github.com/acroca/go-symbols", build.Default.SrcDirs()[1], build.IgnoreVendor)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v, %v\n", pkg, pkg.IsCommand())
}
