// This package tests Exported and Local.
package scope

import (
	// _: local
	_ "fmt"

	// a: local
	s "strings"
)

// Comments indicate the desired scope. The default is unexported and
// nonlocal. Exported and local defs are marked as such.

// a: (unexported)
var a bool

// A: exported
var A bool

// fn
func fn() {
	// b: local
	var b bool
	_ = b

	// T0: local
	// T0.f: local
	type T0 struct{ f string }

	// x: local
	for _, x := range []int{} {
		// y: local
		y := x
		_ = y

		// Y: local
		var Y int
		_ = Y

		// T1: local
		// T1.f: local
		type T1 struct{ f string }
	}
}

// T2: exported
type T2 struct{}

// F: exported
// a: local
// b: local
func (T2) F(a string) (b string) {
	return
}

// g: exported type, unexported method
func (T2) g() {
	return
}

// Unexported type, exported methods

type t3 struct{}

// F: exported
func (t3) F() {}

// g: unexported
func (t3) g() {}

func init() {
	_ = s.Replace
}
