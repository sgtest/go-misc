// Package go_subrepo_import tests that srclib-go properly resolves imports of
// Go subrepositories (code.google.com/p/go.*).
package go_subrepo_import

import (
	"code.google.com/p/go.net/ipv6"
	"code.google.com/p/go.tools/go/types"
)

func Foo() {
	c := ipv6.Conn{}
	_ = c

	t := types.Config{}
	_ = t
}
