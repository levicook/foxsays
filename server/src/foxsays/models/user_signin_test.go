package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserSignin_Errors(t *testing.T) {
	us := UserSignin{}

	errs := us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 2)
	assert.Equal(t, errs["email"], "is required")
	assert.Equal(t, errs["password"], "is required")

	us.Email = "foobar"
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 2)
	assert.Equal(t, errs["email"], "is invalid")
	assert.Equal(t, errs["password"], "is required")

	us.Email = "bart@gmail.com"
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 1)
	assert.Equal(t, errs["password"], "is required")

	us.Password = " "
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 1)
	assert.Equal(t, errs["password"], "is required")

	us.Password = "secret"
	errs = us.Errors()
	assert.Empty(t, errs)

}
