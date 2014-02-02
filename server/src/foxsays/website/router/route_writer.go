package router

import "net/http"

type routeWriter struct {
	writeCalled bool
	http.ResponseWriter
}

func (w *routeWriter) Write(b []byte) (int, error) {
	w.writeCalled = true
	return w.ResponseWriter.Write(b)
}
