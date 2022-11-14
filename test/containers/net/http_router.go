package net

import "net/http"

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(s *Substance) {
	key := s.Method + "-" + s.Path
	if handler, ok := r.handlers[key]; ok {
		handler(s)
	} else {
		s.String(http.StatusNotFound, "404 NOT FOUND: %s\n", s.Path)
	}
}
