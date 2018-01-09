//go:binary-only-package

package struct

import (
	"fmt"
)

type Model struct {

}

func (m *Model) String() string {
	return "abc"
}

type GirlModel struct {
	Model
}

func (gm *GirlModel) String() string {
	return fmt.Sprintln("abc")
}