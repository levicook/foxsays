package models

import (
	"time"
	"foxsays/mime"
)

type (
	Image struct {
		Id          FileId `bson:"id"  json:"id"`
		ContentType string `bson:"ct"  json:"contentType"`
		FileName    string `bson:"nm"  json:"fileName"`
		Size        int64  `bson:"sz"  json:"size"`

		CreatedAt time.Time `bson:"cat" json:"createdAt"`
		CreatedBy UserId    `bson:"cby" json:"createdBy"`
	}
)

func (o Image) Errors() Errors {
	errors := make(Errors)

	if o.Id == "" {
		errors["id"] = "is required"
	}

	if !mime.IsImage(o.ContentType) {
		errors["contentType"] = "is invalid"
	}

	if o.CreatedAt.IsZero() {
		errors["createdAt"] = "is required"
	}

	if o.CreatedBy == "" {
		errors["createdBy"] = "is required"
	}

	if o.Size == 0 {
		errors["size"] = "is required"
	} else if o.Size < 0 {
		errors["size"] = "is invalid"
	}

	return errors
}
