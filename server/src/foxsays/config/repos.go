package config

import (
	"foxsays/repos"

	"labix.org/v2/mgo"
)

type repositories struct {
	session *mgo.Session
}

func (r repositories) NewFileRepo() repos.FileRepo {
	return repos.NewFileRepo(r.openDatabase())
}

func (r repositories) NewImageRepo() repos.ImageRepo {
	return repos.NewImageRepo(r.openDatabase())
}

func (r repositories) NewUserRepo() repos.UserRepo {
	return repos.NewUserRepo(r.openDatabase(), pword.engine())
}

func (r *repositories) Open() (err error) {
	r.session, err = mgo.Dial(mongo.Dial)
	return
}

func (r repositories) Close() {
	if r.session != nil {
		r.session.Close()
	}
}

func (r repositories) openDatabase() *mgo.Database {
	return r.session.DB(mongo.Database)
}
