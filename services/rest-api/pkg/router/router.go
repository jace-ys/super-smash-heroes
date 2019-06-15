package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

type routeHandler func(w http.ResponseWriter, r *http.Request)

func (r *Router) Get(path string, handler routeHandler) {
	r.HandleFunc(path, handler).Methods("GET")
}

func (r *Router) Post(path string, handler routeHandler) {
	r.HandleFunc(path, handler).Methods("POST")
}

func (r *Router) Put(path string, handler routeHandler) {
	r.HandleFunc(path, handler).Methods("PUT")
}

func (r *Router) Delete(path string, handler routeHandler) {
	r.HandleFunc(path, handler).Methods("DELETE")
}
