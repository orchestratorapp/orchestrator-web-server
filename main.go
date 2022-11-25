package main

import (
	"log"
	"net/http"
	middleware "orchestrator_server/middleware"
	server "orchestrator_server/server"
)

func main() {
	s := server.BuildServer()
	s.Handle(http.MethodGet, "/",
		s.AddMiddleware(server.HandleRoot, middleware.RequestLogger()),
	)
	s.Handle(http.MethodPost, "/post",
		s.AddMiddleware(server.PostRequest, middleware.RequestLogger()),
	)
	err := s.Listen()
	if err != nil {
		log.Fatalf("\033[41m FATAL \033[0m %v", err)
	}
}
