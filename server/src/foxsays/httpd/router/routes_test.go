package router

import (
	"testing"
)

func Test_Routes(t *testing.T) {
	nameCounter := make(map[string]int)

	for _, route := range routes {
		nameCounter[route.name]++

		if nameCounter[route.name] > 1 {
			t.Fatalf("duplicate route name: %q", route.name)
		}
	}
}
