package models

type UserSignin struct {
	Email    EmailAddress `json:"email"`
	Password string       `json:"password"`
}

func (o UserSignin) Errors() Errors {
	errors := make(Errors)

	{ // validate email
		email := o.Email

		if email.Blank() {
			errors["email"] = "is required"
		} else if email.Invalid() {
			errors["email"] = "is invalid"
		}
	}

	{ // validate password
		password := o.Password
		if trimSpace(password) == "" {
			errors["password"] = "is required"
		}
	}

	return errors
}
