package repos

import (
	"foxsays/models"

	"labix.org/v2/mgo/bson"
)

func fileOid(id models.FileId) (oid bson.ObjectId) {
	if sid := string(id); bson.IsObjectIdHex(sid) {
		oid = bson.ObjectIdHex(sid)
	}
	return
}
