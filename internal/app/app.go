package app

import (
	"log"
	"net/http"

	"github.com/vishnumohanan404/forminit/service/user"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(router)
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}
	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}
