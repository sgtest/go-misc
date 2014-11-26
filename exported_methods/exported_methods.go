// Package exported_methods tests that srclib-go exports all
// capitalized methods, even if the type that the method is on is
// unexported.
package exported_methods

type unexported struct{}

func (u unexported) ExportedMethod() {
}
