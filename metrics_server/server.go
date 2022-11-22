package metrics_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"regexp"

	"github.com/guiln/prom-gateway/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type metricsServer struct {
	proto.UnimplementedMetricsServer
	lggr         *log.Logger
	promRegistry *prometheus.Registry
	counters     map[string]prometheus.Counter
}

func new(lggr *log.Logger) (*metricsServer, *prometheus.Registry) {
	r := prometheus.NewRegistry()
	r.MustRegister(collectors.NewBuildInfoCollector())
	r.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))

	return &metricsServer{
		lggr:         lggr,
		promRegistry: r,
		counters:     make(map[string]prometheus.Counter),
	}, r
}

func (ms *metricsServer) IncrementCounter(ctx context.Context, req *proto.IncrementCounterRequest) (*proto.IncrementCounterResponse, error) {

	_, exists := ms.counters[req.CounterName]
	if !exists {
		ms.counters[req.CounterName] = promauto.With(ms.promRegistry).NewCounter(prometheus.CounterOpts{
			Name: req.CounterName,
			Help: *req.CounterHelp,
		})
	}

	ms.counters[req.CounterName].Inc()

	return &proto.IncrementCounterResponse{
		Result: proto.IncrementCounterResponse_SUCCESS,
	}, nil
}

// Metrics Grpc server

type MetricsGrpcServer struct {
	PrometheusRegistry *prometheus.Registry
	grpcServer         *grpc.Server
	metricsServer      *metricsServer
	lggr               *log.Logger
}

func NewMetricsGrpcServer(lggr *log.Logger) *MetricsGrpcServer {
	mServer, r := new(lggr)
	return &MetricsGrpcServer{
		PrometheusRegistry: r,
		grpcServer:         nil,
		metricsServer:      mServer,
		lggr:               lggr,
	}
}

func (s *MetricsGrpcServer) StopServer() error {
	s.lggr.Println("Stoping server...")
	if s.grpcServer != nil {
		s.grpcServer.Stop()
		return nil
	}
	return fmt.Errorf("cannot stop a server that has not been started.")
}

func (s *MetricsGrpcServer) RunGrpcServer(ctx context.Context, tcpAddress string) error {
	s.grpcServer = grpc.NewServer()
	defer func() { s.grpcServer = nil }()

	go func() {
		select {
		case <-ctx.Done():
			s.lggr.Print("Stopping gRPC server...")
			s.grpcServer.Stop()
		}
	}()

	// Regiter service
	proto.RegisterMetricsServer(s.grpcServer, s.metricsServer)

	// Register the reflection service that allows client to determine the
	// methods for this gRPC service
	reflection.Register(s.grpcServer)

	listener, err := net.Listen("tcp", tcpAddress)
	if err != nil {
		return err
	}

	err = s.grpcServer.Serve(listener)
	if err != nil {
		s.lggr.Fatal(err)
		return err
	}

	s.lggr.Print("terminating server...")

	return nil
}
