
package parser

import (
	"fmt"
)

type action interface {
	act()
	String() string
}

type (
	accept bool
	shift  int // value is next state index
	reduce int // value is production index
)

func (this accept) act() {}
func (this shift) act()  {}
func (this reduce) act() {}

func (this accept) Equal(that action) bool {
	if _, ok := that.(accept); ok {
		return true
	}
	return false
}

func (this reduce) Equal(that action) bool {
	that1, ok := that.(reduce)
	if !ok {
		return false
	}
	return this == that1
}

func (this shift) Equal(that action) bool {
	that1, ok := that.(shift)
	if !ok {
		return false
	}
	return this == that1
}

func (this accept) String() string { return "accept(0)" }
func (this shift) String() string  { return fmt.Sprintf("shift:%d", this) }
func (this reduce) String() string {
	return fmt.Sprintf("reduce:%d(%s)", this, productionsTable[this].String)
}
