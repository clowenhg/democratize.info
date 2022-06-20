package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)

	go run(ctx, done)

	<-sigs
	cancel()
	<-done
}

func run(ctx context.Context, done chan error) error {
	return nil
}
