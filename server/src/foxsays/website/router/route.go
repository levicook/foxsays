package router

import (
	"foxsays/log"
	"net/http"
	"time"
)

type filterSet []http.HandlerFunc

type route struct {
	name    string
	method  string
	path    string
	filters filterSet
}

func (rt route) logFailure(w http.ResponseWriter, r *http.Request, duration time.Duration, err interface{}) {
	log.Printf(
		"%s %s | %s %s | failure: %v",
		r.Method, r.RequestURI,
		rt.name, duration,
		err,
	)
}

func (rt route) logVictory(w http.ResponseWriter, r *http.Request, duration time.Duration) {
	log.Printf(
		"%s %s | %s %s",
		r.Method, r.RequestURI,
		rt.name, duration,
	)
}

func (rt route) ensureResponseAndLogResults(w http.ResponseWriter, r *http.Request, start time.Time) {
	if err := recover(); err != nil {
		rt.logFailure(w, r, time.Since(start), err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		rt.logVictory(w, r, time.Since(start))
	}
}

func (rt route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer rt.ensureResponseAndLogResults(w, r, time.Now())

	routeWriter := new(routeWriter)
	routeWriter.ResponseWriter = w

	for _, filter := range rt.filters {
		filter.ServeHTTP(routeWriter, r)
		if routeWriter.writeCalled {
			return
		}
	}
}
