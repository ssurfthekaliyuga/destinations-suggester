package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"
)

type Server struct {
	srv  *echo.Echo
	conf *Config
}

func New(srv *echo.Echo, conf *Config) *Server {
	return &Server{
		srv:  srv,
		conf: conf,
	}
}

func (s *Server) Start(_ context.Context, errsCh chan<- error) {
	go func() {
		err := s.srv.StartServer(&http.Server{
			Addr:              net.JoinHostPort(s.conf.Host, s.conf.Port),
			ReadTimeout:       s.conf.ReadTimeout,
			ReadHeaderTimeout: s.conf.ReadHeaderTimeout,
			WriteTimeout:      s.conf.WriteTimeout,
			IdleTimeout:       s.conf.IdleTimeout,
			MaxHeaderBytes:    s.conf.MaxHeaderBytes,
		})
		if err != nil {
			errsCh <- fmt.Errorf("cannot start echo server: %w", err)
		}
	}()
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("cannot shutdown echo server: %w", err)
	}

	return nil
}
