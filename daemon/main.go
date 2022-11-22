package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/guiln/prom-gateway/app"
)

func main() {
	// Get daemon configuration
	daemonCfg, nil := getDaemonConfig()

	// Config contexts
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Config metrics exporter application
	lggr := log.New(os.Stdout, "[METRICS EXPORTER]: ", 0)
	config := &app.MetricsExporterAppConfig{
		IncommingMetricsHandlerGrpcAddress: fmt.Sprintf("%s:%d", daemonCfg.MetricsHandlerAddress, daemonCfg.MetricsHandlerPort),
		PrometheusMetricsEndpointAddress:   fmt.Sprintf("%s:%d", daemonCfg.PrometheusMetricsAddress, daemonCfg.PrometheusMetricsPort),
	}

	// Creates metrics application
	metricsExporterApplication := app.New(lggr, config)

	shouldLoadApplication := true

	// Handle signals
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)

	// Common shutdown function
	shutDown := func(message string) {
		shouldLoadApplication = false
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
					if err := metricsExporterApplication.Terminate(); err != nil {
						lggr.Fatal(err)
						return
					}
					lggr.Printf("cancelling application...")
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

	lggr.Printf("PID: %d", os.Getpid())

	for shouldLoadApplication {
		if err := metricsExporterApplication.Run(ctx); err != nil {
			shutDown("Error while running metrics exporter application")
			return
		}
	}

}

type daemonConfig struct {
	MetricsHandlerPort       int    `json:"metrics_handler_port"`
	MetricsHandlerAddress    string `json:"metrics_handler_address"`
	PrometheusMetricsPort    int    `json:"prometheus_metrics_port"`
	PrometheusMetricsAddress string `json:"prometheus_metrics_address"`
}

func getDaemonConfig() (*daemonConfig, error) {

	configFile := flag.String("config-file", "", "configuration file. Inline configuration has precedence over file configuration")

	metricsHandlerAddress := flag.String("metrics-handler-address", "", "metrics handler address exposed by gRPC")
	prometheusMetricsAddress := flag.String("prometheus-metrics-address", "", "address to be accessed by prometheus scraper")

	metricsHandlerPort := flag.Int("metrics-handler-port", 0, "metrics handler port exposed by gRPC")
	prometheusMetricsPort := flag.Int("prometheus-metrics-port", 0, "port to be accessed by prometheus scraper")

	flag.Parse()

	if *metricsHandlerAddress == "" || *metricsHandlerPort == 0 || *prometheusMetricsAddress == "" || *prometheusMetricsPort == 0 {
		if *configFile == "" {
			return nil,
				fmt.Errorf("Insuficient parameters was provided.")
		}
		return loadConfigFromFile(*configFile)
	}
	return &daemonConfig{
		MetricsHandlerPort:       *metricsHandlerPort,
		MetricsHandlerAddress:    *metricsHandlerAddress,
		PrometheusMetricsPort:    *prometheusMetricsPort,
		PrometheusMetricsAddress: *prometheusMetricsAddress,
	}, nil
}

func loadConfigFromFile(configFile string) (*daemonConfig, error) {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var payload daemonConfig
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
