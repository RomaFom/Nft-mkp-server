package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	httpserver *http.Server
	response   *http.ResponseWriter
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpserver = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logrus.Printf("Server Running on %s", viper.GetString("port"))
	return s.httpserver.ListenAndServe()
}

func (s *Server) EnableCors() {
	(*s.response).Header().Set("Access-Control-Allow-Origin", "*")
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpserver.Shutdown(ctx)
}
