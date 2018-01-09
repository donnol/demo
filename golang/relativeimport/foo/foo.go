package foo

import (
	"log"

	"demo/relativeimport/foo/bar"
)

func init() {
	log.Println(bar.Bar)
}
