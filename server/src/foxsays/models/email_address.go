package models

import "regexp"

type (
	EmailAddresses []EmailAddress
	EmailAddress   string
)

var validator = regexp.MustCompile(`^([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})$`)

func (a EmailAddress) String() string {
	return string(a.Normalize())
}

func (a EmailAddress) Valid() bool {
	return validator.MatchString(a.String())
}

func (a EmailAddress) Invalid() bool {
	return !a.Valid()
}

func (a EmailAddress) Blank() bool   { return a == "" }
func (a EmailAddress) Present() bool { return !a.Blank() }

func (a EmailAddress) Normalize() EmailAddress {
	return EmailAddress(toLower(trimSpace(string(a))))
}

func (addresses EmailAddresses) Normalize() EmailAddresses {
	results := make(EmailAddresses, len(addresses), len(addresses))

	for i, address := range addresses {
		results[i] = address.Normalize()
	}

	return results
}
