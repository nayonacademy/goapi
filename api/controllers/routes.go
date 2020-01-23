package controllers

import "github.com/nayonacademy/goapi/api/middlewares"

// InitialRoutes - Call all routes here
func (s *Server) InitialRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
}
