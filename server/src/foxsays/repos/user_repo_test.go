package repos

import (
	"foxsays/models"
	"foxsays/password"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserRepo(t *testing.T) {
	setupTestDB()
	defer teardownDB()

	e := password.NewTestEngine()
	r := NewUserRepo(testDB, e)
	s := models.UserSignup{
		FirstName: "Bart",
		LastName:  "Simpson",
		Email:     "bart@gmail.com",
		Password:  "super-sekret",
	}
	d := e.Digest(s.Password)

	u1, errs1, err1 := r.CreateBySignup(s)
	assert.Equal(t, u1.PasswordDigest, d)
	assert.Empty(t, errs1)
	assert.Nil(t, err1)

	u2, errs2, err2 := r.CreateBySignup(s)
	assert.Equal(t, u2.PasswordDigest, d)
	assert.NotEmpty(t, errs2.Present())
	assert.Equal(t, 1, len(errs2))
	assert.Equal(t, "is already registered", errs2["email"])
	assert.Nil(t, err2)

	u3, err3 := r.OneById(u1.Id)
	assert.Equal(t, u1.Id, u3.Id)
	assert.Equal(t, "bart@gmail.com", u3.PrimaryEmail)
	assert.Nil(t, err3)

	si := models.UserSignin{Email: s.Email, Password: s.Password}
	u4, err4 := r.OneBySignin(si)
	assert.Equal(t, "bart@gmail.com", u4.PrimaryEmail)
	assert.Nil(t, err4)

	si = models.UserSignin{Email: "bogus@mail.com", Password: s.Password}
	u5, err5 := r.OneBySignin(si)
	assert.False(t, u5.Present())
	assert.Nil(t, err5)

	si = models.UserSignin{Email: s.Email, Password: "bogus-password"}
	u6, err6 := r.OneBySignin(si)
	assert.False(t, u6.Present())
	assert.Nil(t, err6)
}
