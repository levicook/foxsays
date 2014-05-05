package models

import (
	"io"
	"time"
)

type (
	FileId string

	File interface {
		io.Closer
		io.Reader
		io.Writer

		Id() FileId

		ContentType() string
		SetContentType(string)

		Name() string
		SetName(string)

		UploadDate() time.Time

		Size() int64
	}
)
