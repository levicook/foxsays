package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EmailAddress_Normalize(t *testing.T) {
	expected := EmailAddress("bart@gmail.com")
	observed := EmailAddress(" BART@GMAIL.COM ").Normalize()
	assert.Equal(t, expected, observed)
}

func Test_EmailAddresses_Normalize(t *testing.T) {
	expected := EmailAddresses{
		"bart+test1@gmail.com",
		"bart+test2@gmail.com",
	}

	observed := EmailAddresses{
		"   BART+TEST1@GMAIL.COM",
		"BART+TEST2@GMAIL.COM   ",
	}.Normalize()

	assert.Equal(t, expected, observed)
}
