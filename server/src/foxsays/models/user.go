package models

import (
	"time"
)

type (
	UserId  string
	UserIds []UserId
	Users   []User

	User struct {
		Id UserId `bson:"id" json:"id"`

		FirstName string `bson:"f" json:"firstName"`
		LastName  string `bson:"l" json:"firstName"`

		Password       string `bson:"-"  json:"password"`
		PasswordDigest []byte `bson:"pd" json:"-"`

		CreatedAt time.Time `bson:"cat"  json:"createdAt"`
		UpdatedAt time.Time `bson:"uat"  json:"updatedAt"`
	}
)
