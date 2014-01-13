package router

import "net/http"

var pong = []byte{'p', 'o', 'n', 'g'}

func ping(w http.ResponseWriter, r *http.Request) { w.Write(pong) }
