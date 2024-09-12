package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/westford14/gimmedat/internal/cmd"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cmd.Execute(ctx)
}
