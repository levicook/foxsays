package repos

import (
	"fmt"
	"foxsays/log"
	"foxsays/models"
	"foxsays/password"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func NewUserRepo(db *mgo.Database, pw password.Engine) UserRepo {
	c := db.C("users")

	log.PanicIf(c.EnsureIndexKey("u.cat"))

	log.PanicIf(c.EnsureIndex(mgo.Index{
		Key:    []string{"e"},
		Unique: true,
	}))

	return userRepo{c: c, pw: pw}
}

type UserRepo interface {
	CreateBySignup(models.UserSignup) (models.User, models.Errors, error)

	OneBySignin(models.UserSignin) (models.User, error)
	OneById(models.UserId) (models.User, error)

	SetSuperUser(models.UserId, bool) (models.User, error)
}

type userRepo struct {
	c  *mgo.Collection
	pw password.Engine
}

func (r userRepo) CreateBySignup(o models.UserSignup) (u models.User, errors models.Errors, err error) {
	if errs := o.Errors(); errs.Present() {
		err = fmt.Errorf("invalid signup %q %q", o, errs)
		return
	}

	oid := bson.NewObjectId()
	now := time.Now()

	u = models.User{
		Id:             models.UserId(oid.Hex()),
		PasswordDigest: r.pw.Digest(o.Password),
		PrimaryEmail:   o.Email,
		FirstName:      o.FirstName,
		LastName:       o.LastName,
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}

	err = r.c.Insert(userWrapper{
		Id:             oid,
		User:           u,
		EmailAddresses: u.EmailAddresses().Normalize(),
	})

	if err != nil {
		if le, ok := err.(*mgo.LastError); ok && le.Code == 11000 {
			err = nil
			errors = models.Errors{"email": "is already registered"}
		}
	}

	return
}

func (r userRepo) OneBySignin(o models.UserSignin) (models.User, error) {
	return r.oneByEmailAndPassword(o.Email, o.Password)
}

func (r userRepo) oneByEmailAndPassword(e models.EmailAddress, password string) (o models.User, err error) {
	var w userWrapper
	if err = r.c.Find(m{"e": e}).One(&w); err != nil {
		if err == mgo.ErrNotFound {
			err = nil
		}
		return
	}

	if !r.pw.Equal(r.pw.Digest(password), w.User.PasswordDigest) {
		return
	}

	o = w.User
	return
}

func (r userRepo) OneById(id models.UserId) (o models.User, err error) {
	var w userWrapper
	if err = r.c.FindId(userOid(id)).One(&w); err != nil {
		if err == mgo.ErrNotFound {
			err = nil
		}
		return
	}

	o = w.User
	return
}

func (r userRepo) SetSuperUser(id models.UserId, superUser bool) (o models.User, err error) {
	change := mgo.Change{
		ReturnNew: true,
		Update: m{
			"$set": m{
				"u.su":  superUser,
				"u.uat": time.Now(),
			}}}

	var w userWrapper
	if _, err = r.c.FindId(userOid(id)).Apply(change, &w); err != nil {
		return
	}

	o = w.User
	return
}

type userWrapper struct {
	Id             bson.ObjectId         `bson:"_id"`
	User           models.User           `bson:"u"`
	EmailAddresses models.EmailAddresses `bson:"e"`
}
