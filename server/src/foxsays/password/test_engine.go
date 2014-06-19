package password

import "bytes"

func NewTestEngine() Engine {
	return testEngine{}
}

type testEngine struct{}

func (e testEngine) Digest(pass string) []byte {
	return []byte(pass)
}

func (e testEngine) Equal(pass1, pass2 []byte) bool {
	return bytes.Equal(pass1, pass2)
}
