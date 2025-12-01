package consumers

import (
	"context"
)

type errorsHandler interface {
	Handle(ctx context.Context, msg string, err error)
}
