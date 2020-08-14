package controllers

import (
	"github.com/Lgdev07/superapi/api/middlewares"
)

func (s *Server) InitializeRoutes() {
	s.Router.Use(middlewares.SetContentTypeMiddleware)

	s.Router.HandleFunc("/index", s.Index).Methods("GET")
	s.Router.HandleFunc("/supers", s.CreateSuper).Methods("POST")
	s.Router.HandleFunc("/supers", s.ListSupers).Methods("GET")
	s.Router.HandleFunc("/supers/{id:[0-9]+}", s.Index).Methods("DELETE")
}
