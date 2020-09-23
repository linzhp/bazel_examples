package recompile_test

import (
	"github.com/linzhp/bazel_examples/recompile"
	_ "github.com/linzhp/bazel_examples/recompile/bar"
)

func NewThing(name string, size int) recompile.Thing {
	return recompile.Thing{"Uber", 1}
}
