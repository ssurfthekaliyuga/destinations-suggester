package server

import "go.uber.org/fx"

var Fx = fx.Module("http_server", fx.Provide())
