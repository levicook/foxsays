package detect

import "testing"

func Test_String(t *testing.T) {
	s := String("a", "b")
	if s != "a" {
		t.Fatalf("%#v", s)
	}
}
