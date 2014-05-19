package repos

import (
	"foxsays/models"

	"labix.org/v2/mgo/bson"
)

type m bson.M
type d bson.D

func fileOid(id models.FileId) bson.ObjectId { return safeOid(string(id)) }
func userOid(id models.UserId) bson.ObjectId { return safeOid(string(id)) }

func safeOid(sid string) bson.ObjectId {
	if !bson.IsObjectIdHex(sid) {
		sid = "000000000000000000000000"
	}
	return bson.ObjectIdHex(sid)
}
