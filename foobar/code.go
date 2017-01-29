package foobar

import "github.com/thockin-test/util/bar"
import "github.com/thockin-test/util/foo"

func FooBar() string {
	return foo.Foo() + bar.Bar()
}
