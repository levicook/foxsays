package sessions

import (
	"foxsays/config"
	"foxsays/log"
	"foxsays/models"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
)

func Get(r *http.Request) Session {
	store := config.Httpd.NewSessionStore()
	gs, _ := store.Get(r, config.Httpd.SessionName)

	return &session{r: r, gs: gs}
}

type Session interface {
	Clear(http.ResponseWriter)
	Save(http.ResponseWriter)

	RealUser() models.User
	RealUserId() models.UserId
	SetRealUserId(models.UserId)

	EffectiveUser() models.User
	EffectiveUserId() models.UserId
	SetEffectiveUserId(models.UserId)
}

const (
	effectiveUserIdKey = "euid"
	realUserIdKey      = "ruid"
)

type session struct {
	r  *http.Request
	gs *sessions.Session

	realUser     models.User
	realUserOnce sync.Once

	effectiveUser     models.User
	effectiveUserOnce sync.Once
}

func (s *session) Clear(w http.ResponseWriter) {
	delete(s.gs.Values, effectiveUserIdKey)
	delete(s.gs.Values, realUserIdKey)
	s.Save(w)
}

func (s *session) RealUserId() (id models.UserId) { return s.userId(realUserIdKey) }
func (s *session) SetRealUserId(id models.UserId) { s.gs.Values[realUserIdKey] = string(id) }

func (s *session) EffectiveUserId() models.UserId      { return s.userId(effectiveUserIdKey) }
func (s *session) SetEffectiveUserId(id models.UserId) { s.gs.Values[effectiveUserIdKey] = string(id) }

func (s *session) RealUser() models.User {
	s.realUserOnce.Do(func() {
		userRepo := config.Repos.NewUserRepo()
		user, err := userRepo.OneById(s.RealUserId())
		log.PanicIf(err)
		s.realUser = user
	})
	return s.realUser
}

func (s *session) EffectiveUser() models.User {
	s.effectiveUserOnce.Do(func() {
		userRepo := config.Repos.NewUserRepo()
		user, err := userRepo.OneById(s.EffectiveUserId())
		log.PanicIf(err)
		s.effectiveUser = user
	})
	return s.realUser
}

func (s *session) userId(key string) (id models.UserId) {
	if v, ok := s.gs.Values[key]; ok {
		id = models.UserId(v.(string))
	}
	return
}

func (s *session) Save(w http.ResponseWriter) {
	log.FatalIf(s.gs.Save(s.r, w))
}
