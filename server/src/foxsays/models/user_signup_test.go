package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserSignup(t *testing.T) {
	us := UserSignup{}

	errs := us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 4)
	assert.Equal(t, errs["firstName"], "is required")
	assert.Equal(t, errs["lastName"], "is required")
	assert.Equal(t, errs["email"], "is required")
	assert.Equal(t, errs["password"], "is required")

	us.Email = "foobar"
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 4)
	assert.Equal(t, errs["firstName"], "is required")
	assert.Equal(t, errs["lastName"], "is required")
	assert.Equal(t, errs["email"], "is invalid")
	assert.Equal(t, errs["password"], "is required")

	us.Email = "bart@gmail.com"
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 3)
	assert.Equal(t, errs["firstName"], "is required")
	assert.Equal(t, errs["lastName"], "is required")
	assert.Equal(t, errs["password"], "is required")

	us.Password = " "
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 3)
	assert.Equal(t, errs["firstName"], "is required")
	assert.Equal(t, errs["lastName"], "is required")
	assert.Equal(t, errs["password"], "is required")

	us.Password = "secret"
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 2)
	assert.Equal(t, errs["firstName"], "is required")
	assert.Equal(t, errs["lastName"], "is required")

	us.FirstName = " "
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 2)
	assert.Equal(t, errs["firstName"], "is required")
	assert.Equal(t, errs["lastName"], "is required")

	us.FirstName = "Levi"
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 1)
	assert.Equal(t, errs["lastName"], "is required")

	us.LastName = " "
	errs = us.Errors()
	assert.NotEmpty(t, errs)
	assert.Equal(t, len(errs), 1)
	assert.Equal(t, errs["lastName"], "is required")

	us.LastName = "Cook"
	errs = us.Errors()
	assert.Empty(t, errs)
}
