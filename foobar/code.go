package foobar

import (
	"strings"

	"github.com/thockin-test/util/bar"
	"github.com/thockin-test/util/foo"
)

func FooBar() string {
	return strings.ToUpper(foo.Foo() + bar.Bar())
}
