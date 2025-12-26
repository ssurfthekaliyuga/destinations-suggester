package sl

import "log/slog"

func WithComponent(component string) *slog.Logger {
	return slog.With(Component(component))
}
