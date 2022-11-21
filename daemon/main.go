package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/guiln/prom-gateway/app"
)

func main() {

	// Config contexts
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Config metrics exporter application
	lggr := log.New(os.Stdout, "[METRICS EXPORTER]: ", 0)
	config := &app.MetricsExporterAppConfig{
		IncommingMetricsHandlerGrpcAddress: "0.0.0.0:50051",
		PrometheusMetricsEndpointAddress:   "0.0.0.0:8080",
	}
	metricsExporterApplication := app.New(lggr, config)

	// Handle signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)

	// Common shutdown function
	shutDown := func(message string) {
		lggr.Printf(message)
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
					shutDown("Got interrupeted. Shutting down...")
					return
				case syscall.SIGHUP:
					lggr.Printf("Reloading...")
					if err := metricsExporterApplication.Reload(); err != nil {
						shutDown("Failed to reload application.")
						return
					}
					lggr.Printf("Applicattion successful reloaded.")
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

	if err := metricsExporterApplication.Run(ctx); err != nil {
		shutDown("Error while running metrics exporter application")
		return
	}
}
