package controllers

import (
	"github.com/febrielven/go_backend_webcompro/api/middlewares"
)

func (s *Server) initallizeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/all", middlewares.SetMiddlewareJSON(s.GetAll)).Methods("GET")
	s.Router.HandleFunc("/save", middlewares.SetMiddlewareJSON(s.Create)).Methods("POST")
}
