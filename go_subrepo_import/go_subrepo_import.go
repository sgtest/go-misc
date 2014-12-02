// Package go_subrepo_import tests that srclib-go properly resolves imports of
// Go subrepositories (golang.org/x/*).
package go_subrepo_import

import (
	"golang.org/x/net/ipv6"
	"golang.org/x/tools/go/types"
)

func Foo() {
	c := ipv6.Conn{}
	_ = c

	t := types.Config{}
	_ = t
}
