package httpx

import (
	"net/http"
	"time"
)

type Server struct {
	host   string
	port   string
	server *http.Server
}

func NewServer(host, port string, router *Router) *Server {
	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: router.engine,

		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{
		host:   host,
		port:   port,
		server: server,
	}
}

func (s *Server) Address() string {
	return s.host + ":" + s.port
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.server.Close()
}
