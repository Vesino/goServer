package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

// Create a new Server instance
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

func (s *Server) AddMidleWare(hf http.HandlerFunc, middlewares ...MiddleWare) http.HandlerFunc {
	for _, middleware := range middlewares {
		hf = middleware(hf)
	}
	return hf
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
