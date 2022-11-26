package server

import (
	"fmt"
	"net/http"

	"github.com/orchestratorapp/orchestrator-web-server/config"
	"github.com/orchestratorapp/orchestrator-web-server/middleware"
)

// The web server
type Server struct {
	name    string
	port    string
	router  *Router
	profile *config.ProfileConfig
}

// Create a Server instance
func BuildServer() *Server {
	config, profileConfig := config.LoadConfig()
	return &Server{
		name:    config.Orchestrator.Server.AppName,
		port:    config.Orchestrator.Server.Port,
		router:  BuildRouter(),
		profile: profileConfig,
	}
}

// Listen for new requests
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	fmt.Printf(
		"\033[35m%s\033[0m is alive and listening on port \033[32m%s\033[0m\n",
		s.name, s.port,
	)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

// Assigns the provided handler function to the provided path
func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exists := s.router.rules[path]
	if !exists {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

// Adds and activates the provided middlewares
func (s *Server) AddMiddleware(
	f http.HandlerFunc,
	middlewares ...middleware.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
