package models

func NewImage(file File, createdBy UserId) Image {
	return Image{
		Id:          file.Id(),
		ContentType: file.ContentType(),
		FileName:    file.Name(),
		Size:        file.Size(),
		CreatedAt:   file.UploadDate(),
		CreatedBy:   createdBy,
	}
}
