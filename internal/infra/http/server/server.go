package server

import "github.com/labstack/echo/v4"

type Server struct {
	echo *echo.Echo
}

func New() *Server {
	srv := echo.New()

	srv.Use()

	return &Server{
		echo: srv,
	}
}

func (s *Server) Start() {
	s.echo.Start("8080")
}
