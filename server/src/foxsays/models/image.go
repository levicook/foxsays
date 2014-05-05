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
		errors["Id"] = "is required"
	}

	if !mime.IsImage(o.ContentType) {
		errors["ContentType"] = "is invalid"
	}

	if o.CreatedAt.IsZero() {
		errors["CreatedAt"] = "is required"
	}

	if o.CreatedBy == "" {
		errors["CreatedBy"] = "is required"
	}

	if o.Size == 0 {
		errors["Size"] = "is required"
	} else if o.Size < 0 {
		errors["Size"] = "is invalid"
	}

	return errors
}
