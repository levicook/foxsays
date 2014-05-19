package password

type Engine interface {
	Digest(string) []byte
	Equal(pass1, pass2 []byte) bool
}
