package app

import (
	"context"
	"log"
	"net/http"

	"github.com/guiln/prom-gateway/metrics_server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsExporterApp struct {
	lggr   *log.Logger
	config *MetricsExporterAppConfig
}

type MetricsExporterAppConfig struct {
	// IncommingMetricsHandlerGrpcAddress
	//
	// Expected: host:port
	IncommingMetricsHandlerGrpcAddress string
	// PrometheusMetricsEndpointAddress
	//
	// Expected: host:port
	PrometheusMetricsEndpointAddress string
}

func New(lggr *log.Logger, config *MetricsExporterAppConfig) *MetricsExporterApp {
	return &MetricsExporterApp{
		lggr:   lggr,
		config: config,
	}
}

// Run
// Main loop that runs the daemon.
func (app *MetricsExporterApp) Run(ctx context.Context) error {
	app.lggr.Println("starting application...")
	ctx, cancelChild := context.WithCancel(ctx)
	defer cancelChild()

	errChannel := make(chan error)

	app.lggr.Printf("Starting incomming metrics server at: %s\n", app.config.IncommingMetricsHandlerGrpcAddress)
	go func() {
		if err := metrics_server.RunGrpcServer(ctx, app.config.IncommingMetricsHandlerGrpcAddress, app.lggr); err != nil {
			errChannel <- err
		}
	}()

	app.lggr.Printf("Starting prometheus metrics endpoint at: %s\n", app.config.PrometheusMetricsEndpointAddress)
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(app.config.PrometheusMetricsEndpointAddress, nil); err != nil {
			errChannel <- err
		}
	}()

	app.lggr.Println("application started!")
	for {
		select {
		case err := <-errChannel:
			app.lggr.Fatalf("Error in child applications...")
			cancelChild()
			return err
		case <-ctx.Done():
			app.lggr.Println("Exiting application...")
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

func startGrpcServer(ctx context.Context) {

}
