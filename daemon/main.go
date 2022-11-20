package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)

	shutDown := func(message string) {
		log.Printf(message)
		cancel()
		os.Exit(1)
		return
	}

	// OS signal handler.
	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGINT, syscall.SIGTERM:
					shutDown("Received term/int signal. Shutting down...")
					return
				case os.Interrupt:
					shutDown("Gor interrupeted. Shutting down...")
					return
				case syscall.SIGHUP:
					// TODO: reload application
					log.Printf("Reloading...")
				}
			case <-ctx.Done():
				shutDown("Done...")
				return
			}
		}
	}()

	// Run app.
	defer func() {
		signal.Stop(signalChan)
		close(signalChan)
		cancel()
	}()
}
