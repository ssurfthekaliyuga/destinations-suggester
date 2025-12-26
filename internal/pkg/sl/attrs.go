package sl

import "log/slog"

const (
	KeyError     = "error"
	KeyComponent = "component"
)

func Error(err error) slog.Attr {
	return slog.String(KeyError, err.Error())
}

func Component(component string) slog.Attr {
	return slog.String(KeyComponent, component)
}
