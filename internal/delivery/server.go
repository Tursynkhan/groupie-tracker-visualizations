package delivery

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer http.Server
}

func (s *Server) ServerRun(addr string, handler http.Handler) error {
	s.httpServer = http.Server{
		Addr:         addr,
		Handler:      handler,
		WriteTimeout: time.Second * 3,
		ReadTimeout:  time.Second * 3,
	}
	log.Printf("Запуск веб-сервера на http://127.0.0.1%s", addr)
	return s.httpServer.ListenAndServe()
}
