package route

import "foxsays/models"

type Vars map[string]string

func (v Vars) UserId(k string) models.UserId {
	return models.UserId(v[k])
}

func (v Vars) FileId(k string) models.FileId {
	return models.FileId(v[k])
}
