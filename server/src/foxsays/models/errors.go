package models

type Errors map[string]string

func (e Errors) Present() bool {
	return len(e) > 0
}
