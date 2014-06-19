package repos

import (
	"foxsays/log"
	"testing"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func Test_safeOid(t *testing.T) {
	oid := safeOid("")
	if oid != bson.ObjectIdHex("000000000000000000000000") {
		t.Fatalf("bad oid: %q", oid)
	}
}

var testDB *mgo.Database

func setupTestDB() {
	s, err := mgo.Dial("127.0.0.1")
	log.PanicIf(err)

	db := s.DB("foxsays-test")
	log.PanicIf(db.DropDatabase())

	testDB = db
}

func teardownDB() {
	if testDB != nil && testDB.Session != nil {
		testDB.Session.Close()
	}
}
