package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/guiln/prom-gateway/metrics_server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
)

type MetricsExporterApp struct {
	lggr           *log.Logger
	config         *MetricsExporterAppConfig
	cancelMainLoop context.CancelFunc
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

	ctx, cancel := context.WithCancel(ctx)
	app.cancelMainLoop = cancel
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	app.lggr.Printf("Starting incomming metrics server at: %s\n", app.config.IncommingMetricsHandlerGrpcAddress)
	metricsGrpcServer := metrics_server.NewMetricsGrpcServer(app.lggr)

	g.Go(func() error {
		return metricsGrpcServer.RunGrpcServer(ctx, app.config.IncommingMetricsHandlerGrpcAddress)
	})

	app.lggr.Printf("Starting prometheus metrics endpoint at: %s\n", app.config.PrometheusMetricsEndpointAddress)
	g.Go(func() error {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.HandlerFor(
			metricsGrpcServer.PrometheusRegistry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		))

		httpServer := &http.Server{Addr: app.config.PrometheusMetricsEndpointAddress, Handler: mux}

		go func() {
			select {
			case <-ctx.Done():
				app.lggr.Println("Shutting down metrics endpoint...")
				if err := httpServer.Close(); err != nil {
					app.lggr.Fatal(err)
				}
			}
		}()

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.lggr.Println("Error while running metrics endpoint server")
			app.lggr.Fatal(err)
			return err
		}

		return nil
	})

	app.lggr.Println("application started!")

	return g.Wait()
}

// Terminates main loop by ending it context
func (app *MetricsExporterApp) Terminate() error {
	app.lggr.Println("Terminating application loop...")
	if app.cancelMainLoop == nil {
		return fmt.Errorf("No application started yet!")
	}
	app.cancelMainLoop()
	return nil
}
