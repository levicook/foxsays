package config

import (
	"foxsays/log"

	"labix.org/v2/mgo"
)

var Mongo mongo

type mongo struct {
	Dial     string
	Database string

	session *mgo.Session
}

func (m *mongo) Open() {
	if m.session == nil {
		s, e := mgo.Dial(m.Dial)
		log.FatalIf(e)
		m.session = s
	}
}

func (m *mongo) Close() {
	if m.session != nil {
		m.session.Close()
	}
}
