package controllers

import "backend-go/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT")

	s.Router.HandleFunc("/book", middlewares.SetMiddlewareJSON(s.CreateBook)).Methods("POST")
	s.Router.HandleFunc("/book/lend/{id}", middlewares.SetMiddlewareJSON(s.LendBook)).Methods("PUT")
}