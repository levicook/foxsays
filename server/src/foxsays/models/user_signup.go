package models

type UserSignup struct {
	FirstName string       `bson:"-" json:"firstName"`
	LastName  string       `bson:"-" json:"lastName"`
	Email     EmailAddress `bson:"-" json:"email"`
	Password  string       `bson:"-" json:"password"`
}

func (o UserSignup) Errors() Errors {
	errors := make(Errors)

	{ // validate firstName
		firstName := trimSpace(o.FirstName)
		if firstName == "" {
			errors["firstName"] = "is required"
		}
	}

	{ // validate lastName
		lastName := trimSpace(o.LastName)
		if lastName == "" {
			errors["lastName"] = "is required"
		}
	}

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
		} else if len(password) < 6 {
			errors["password"] = "is too short"
		}
	}

	return errors
}
