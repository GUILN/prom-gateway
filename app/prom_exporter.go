package app

import (
	"context"
	"log"
)

type MetricsExporterApp struct {
	lggr *log.Logger
}

func New(lggr *log.Logger) *MetricsExporterApp {
	return &MetricsExporterApp{
		lggr: lggr,
	}
}

// Run
// Main loop that runs the daemon.
func (app *MetricsExporterApp) Run(ctx context.Context) error {
	app.lggr.Println("starting application...")

	app.lggr.Println("application started!")
	for {
		select {
		case <-ctx.Done():
			return nil
		}
	}
}

// Reload.
//
// Reloads the application. To be called on system's sighup
func (app *MetricsExporterApp) Reload() error {
	app.lggr.Println("Reloading application...")
	return nil
}
