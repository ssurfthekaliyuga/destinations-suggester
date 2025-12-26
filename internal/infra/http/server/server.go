package server

import (
	"context"
	"destinations-suggester/internal/pkg/sl"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net"
	"net/http"
)

type Server struct {
	logger *slog.Logger
	srv    *echo.Echo
	cancel context.CancelCauseFunc
	conf   *Config
}

func New(srv *echo.Echo, cancel context.CancelCauseFunc, conf *Config) *Server {
	return &Server{
		logger: sl.WithComponent("http.server.Server"),
		srv:    srv,
		cancel: cancel,
		conf:   conf,
	}
}

func (s *Server) Start(_ context.Context) error {
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
			s.logger.Error("cannot start http server", sl.Error(err))
			s.cancel(err)
		}

		s.logger.Info("http server started",
			slog.Any("host", s.conf.Host),
			slog.Any("port", s.conf.Port),
		)
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("cannot shutdown echo server: %w", err)
	}

	return nil
}
