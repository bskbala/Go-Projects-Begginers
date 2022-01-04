package examples

import (
	"net/http"

	"google.golang.org/appengine"
)

//Gaehandler Creates http.Handler to run  in  the google app Engine

//Routes
// GET/ping return "pong"
func Gaehandler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		_ = appengine.NewContext(r)
		w.Write([]byte("pong"))
	})
	return m
}
