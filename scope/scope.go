// This package tests Exported and Local.
package scope

// a - Exported: false, Local: false
var a bool

// A - Exported: true, Local: false
var A bool

// fn - Exported: false, Local: false
func fn() {
	// b - Exported: false, Local: true
	var b bool
	_ = b
}
