package route

import (
	"fmt"
	"net/http"
	"time"
	"foxsays/log"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func (rt Route) Signature() string {
	return fmt.Sprintf("%s %s", rt.Method, rt.Path)
}

func (rt Route) logFailure(r *http.Request, duration time.Duration, err interface{}) {
	log.Printf("%s %s | %s %s | error: %v", r.Method, r.RequestURI, rt.Name, duration, err)
}

func (rt Route) logVictory(r *http.Request, duration time.Duration) {
	log.Printf("%s %s | %s %s", r.Method, r.RequestURI, rt.Name, duration)
}

func (rt Route) log(r *http.Request, start time.Time) {
	if err := recover(); err != nil {
		rt.logFailure(r, time.Since(start), err)
		panic(err)
	}
	rt.logVictory(r, time.Since(start))
}

func (rt Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer rt.log(r, time.Now())
	rt.Handler.ServeHTTP(w, r)
}
