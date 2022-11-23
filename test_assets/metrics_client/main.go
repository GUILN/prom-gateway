package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/guiln/prom-gateway/proto"
	"google.golang.org/grpc"
)

func main() {
	lggr := log.New(os.Stdout, "[CLIENT] ", 0)
	lggr.Println("Starting client...")

	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	if *serverAddress == "" {
		panic("Please provide a address argument, like:  --address 0.0.0.0:50051")
	}
	lggr.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		lggr.Fatal("cannot dial server: ", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := run(ctx, conn, lggr); err != nil {
		lggr.Fatal("error while sending request to the server: ", err)
	}
}

func run(ctx context.Context, conn *grpc.ClientConn, lggr *log.Logger) error {
	for {
		var metricName string
		fmt.Println("Metric Name to increment:")
		fmt.Scan(&metricName)

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		metricsClient := proto.NewMetricsClient(conn)
		req := &proto.IncrementCounterRequest{
			CounterName: metricName,
			CounterHelp: new(string),
		}

		response, err := metricsClient.IncrementCounter(ctx, req)
		if err != nil {
			return err
		}

		if response.Result == proto.IncrementCounterResponse_FAILED {
			return fmt.Errorf("Failed to increment")
		}

		lggr.Printf("%s metric successful incremented!", metricName)
	}
}
