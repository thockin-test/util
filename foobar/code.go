package foobar

import (
	"strings"

	"github.com/thockin-test/util/bar"
	"github.com/thockin-test/util/foo"
)

func FooBar() string {
	return strings.ToLower(foo.Foo() + bar.Bar())
}
