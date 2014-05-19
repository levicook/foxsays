package models

import (
	"time"
)

type (
	UserId  string
	UserIds []UserId
	Users   []User

	User struct {
		Id             UserId `bson:"id" json:"id,omitempty"`
		PasswordDigest []byte `bson:"pd" json:"-"`

		PrimaryEmail   EmailAddress `bson:"pe" json:"primaryEmail,omitempty"`
		SecondaryEmail EmailAddress `bson:"se" json:"secondaryEmail,omitempty"`

		FirstName string `bson:"fn" json:"firstName,omitempty"`
		LastName  string `bson:"ln" json:"lastName,omitempty"`

		SuperUser bool `bson:"su" json:"superUser"`

		CreatedAt *time.Time `bson:"cat"  json:"createdAt,omitempty"`
		UpdatedAt *time.Time `bson:"uat"  json:"updatedAt,omitempty"`
	}
)

func (id UserId) Blank() bool { return id == "" }
func (u User) Blank() bool    { return u.Id.Blank() }

func (id UserId) Present() bool { return !id.Blank() }
func (u User) Present() bool    { return !u.Blank() }

func (u User) EmailAddresses() (addresses EmailAddresses) {
	if u.PrimaryEmail.Present() {
		addresses = append(addresses, u.PrimaryEmail)
	}

	if u.SecondaryEmail.Present() {
		addresses = append(addresses, u.PrimaryEmail)
	}

	return
}
