package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(h http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":80",
			Handler:      h,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) Run() {
	s.httpServer.ListenAndServe()
}
