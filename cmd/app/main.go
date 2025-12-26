package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	go func() {
		sig := <-sigCh
		cancel(fmt.Errorf("received signal: %s", sig))
	}()

	<-ctx.Done()
	slog.Info("app shutting down", slog.String("cause", context.Cause(ctx).Error()))
}
