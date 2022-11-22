package metrics_server

import (
	"context"
	"log"
	"net"

	"github.com/guiln/prom-gateway/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type metricsServer struct {
	proto.UnimplementedMetricsServer
	lggr     *log.Logger
	counters map[string]prometheus.Counter
}

func new(lggr *log.Logger) *metricsServer {
	return &metricsServer{
		lggr:     lggr,
		counters: make(map[string]prometheus.Counter),
	}
}

func (ms *metricsServer) IncrementCounter(ctx context.Context, req *proto.IncrementCounterRequest) (*proto.IncrementCounterResponse, error) {
	_, exists := ms.counters[req.CounterName]
	if !exists {
		ms.counters[req.CounterName] = promauto.NewCounter(prometheus.CounterOpts{
			Name: req.CounterName,
			Help: *req.CounterHelp,
		})
	}

	ms.counters[req.CounterName].Inc()

	return &proto.IncrementCounterResponse{
		Result: proto.IncrementCounterResponse_SUCCESS,
	}, nil
}

func RunGrpcServer(ctx context.Context, tcpAddress string, lggr *log.Logger) error {
	metricsServer := new(lggr)
	grpcServer := grpc.NewServer()

	// Regiter service
	proto.RegisterMetricsServer(grpcServer, metricsServer)

	// Register the reflection service that allows client to determine the
	// methods for this gRPC service
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", tcpAddress)
	if err != nil {
		return err
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
