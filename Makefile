
gen_grpc_interface:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/metrics.proto

run_daemon_process:
	go run ./daemon/main.go

run_test_client:
	go run ./metrics_client/main.go --address 0.0.0.0:50051
