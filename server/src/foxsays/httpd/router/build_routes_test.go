package router

import (
	"testing"
)

func Test_buildRoutes_routeNames(t *testing.T) {
	counter := make(map[string]int)

	for _, route := range buildRoutes() {
		name := route.Name

		counter[name]++
		if counter[name] > 1 {
			t.Fatalf("duplicate route name: %q", name)
		}
	}
}

func Test_buildRoutes_routeSignaures(t *testing.T) {
	counter := make(map[string]int)

	for _, route := range buildRoutes() {
		signature := route.Signature()

		counter[signature]++
		if counter[signature] > 1 {
			t.Fatalf("duplicate route signature: %q", signature)
		}
	}
}
