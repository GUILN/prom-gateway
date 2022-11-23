// In order to run the tests contained in this package
// you must start the promgateway configured according
// to the tests configuration.
package test

import (
	"fmt"
	"os"
)

type configEnvVar string

const (
	handlerPort        configEnvVar = "HANDLER_PORT"
	handlerAddress     configEnvVar = "HANDLER_ADDRESS"
	promMetricsPort    configEnvVar = "METRICS_PORT"
	promMetricsAddress configEnvVar = "METRICS_ADDRESS"
)

type testConfig struct {
	MetricsHandlerPort       string `json:"metrics_handler_port"`
	MetricsHandlerAddress    string `json:"metrics_handler_address"`
	PrometheusMetricsPort    string `json:"prometheus_metrics_port"`
	PrometheusMetricsAddress string `json:"prometheus_metrics_address"`
}

func (tc *testConfig) GetFullMetricsAddress() string {
	return fmt.Sprintf("%s:%s", tc.MetricsHandlerAddress, tc.MetricsHandlerPort)
}

func (tc *testConfig) GetFullPromMetricsAddress() string {
	return fmt.Sprintf("%s:%s", tc.PrometheusMetricsAddress, tc.PrometheusMetricsPort)
}

func CreateConfigFromEnvVars() (*testConfig, error) {
	hp := os.Getenv(string(handlerPort))
	ha := os.Getenv(string(handlerAddress))
	pmp := os.Getenv(string(promMetricsPort))
	pma := os.Getenv(string(promMetricsAddress))

	if hp == "" || ha == "" || pmp == "" || pma == "" {
		return nil, fmt.Errorf("Configuration values must be provided")
	}

	return &testConfig{
		MetricsHandlerPort:       hp,
		MetricsHandlerAddress:    ha,
		PrometheusMetricsPort:    pmp,
		PrometheusMetricsAddress: pma,
	}, nil
}
