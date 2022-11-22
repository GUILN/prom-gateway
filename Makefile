
gen_grpc_interface:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/metrics.proto

run_daemon_process:
	go run ./daemon/main.go --metrics-handler-address 0.0.0.0 \
		--metrics-handler-port 50051 \
		--prometheus-metrics-address 0.0.0.0 \
		--prometheus-metrics-port 8080 \

run_test_client:
	go run ./metrics_client/main.go --address 0.0.0.0:50051

## Building
build:
	@echo "Building daemon binary..."
	go build -o bin/promgateway daemon/main.go
	@echo "Finished building daemon binary"
