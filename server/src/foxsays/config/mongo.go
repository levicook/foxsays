package config

import (
	"foxsays/repos"

	"labix.org/v2/mgo"
)

type mongo struct {
	Dial     string `toml:"dial"`
	Database string `toml:"database"`

	session *mgo.Session
}

func (m *mongo) OpenFileRepo() repos.FileRepo {
	return repos.OpenFileRepo(m.openDatabase())
}

func (m *mongo) OpenImageRepo() repos.ImageRepo {
	return repos.OpenImageRepo(m.openDatabase())
}

func (m *mongo) Open() (err error) {
	m.session, err = mgo.Dial(m.Dial)
	return
}

func (m mongo) Close() {
	if m.session != nil {
		m.session.Close()
	}
}

func (m *mongo) openDatabase() *mgo.Database {
	return m.session.DB(m.Database)
}
