package controllers

import (
	"github.com/Lgdev07/superapi/api/middlewares"
)

// InitializeRoutes set the routes and middlewares for the Router
func (s *Server) InitializeRoutes() {
	s.Router.Use(middlewares.SetContentTypeMiddleware)

	s.Router.HandleFunc("/supers", s.CreateSuper).Methods("POST")
	s.Router.HandleFunc("/supers", s.ListSupers).Methods("GET")
	s.Router.HandleFunc("/supers/{id:[a-z0-9-]+}", s.ShowSuper).Methods("GET")
	s.Router.HandleFunc("/supers/{id:[a-z0-9-]+}", s.DeleteSuper).Methods("DELETE")
}
