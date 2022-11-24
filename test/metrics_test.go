package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/guiln/prom-gateway/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

// Common vars
var configuration *testConfig
var metricsClient proto.MetricsClient
var promMetricsEndpoint string

func httpGetBody(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(bodyStr), nil
}

func TestMain(m *testing.M) {
	var err error
	configuration, err = CreateConfigFromEnvVars()
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(configuration.GetFullMetricsAddress(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	metricsClient = proto.NewMetricsClient(conn)

	promMetricsEndpoint = fmt.Sprintf("http://%s/metrics", configuration.GetFullPromMetricsAddress())

	os.Exit(m.Run())
}

func Test_Promgateway_IncrementCounter_IncrementsAndExposes_CounterInConfiguredEndpoints(t *testing.T) {
	myCounter := "my_counter"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Make the call
	req := &proto.IncrementCounterRequest{
		CounterName: myCounter,
		CounterHelp: new(string),
	}

	// Assert response
	response, err := metricsClient.IncrementCounter(ctx, req)
	fmt.Printf("%v\n", err)
	assert.Nil(t, err)
	assert.Equal(t, proto.IncrementCounterResponse_SUCCESS, response.Result)

	// Assert endpoint exposes and updates counter
	exposedMetrics, err := httpGetBody(promMetricsEndpoint)
	assert.Nil(t, err)

	assert.Contains(t, exposedMetrics, fmt.Sprintf("%s 1", myCounter))
}
