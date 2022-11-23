
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
	cp ./promgateway.conf.json bin
	@echo "Finished building daemon binary"

build_darwin:
	@echo "Building daemon binary..."
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/promgateway daemon/main.go
	cp ./promgateway.conf.json bin/darwin/
	@echo "Finished building daemon binary"

build_linux:
	@echo "Building daemon binary..."
	GOOS=linux GOARCH=arm go build -o bin/linux/promgateway daemon/main.go
	cp ./promgateway.conf.json bin/linux/
	@echo "Finished building daemon binary"

build_installer:
	@echo "Building installer..."
	go build -o bin/installer daemon_installer/main.go
	@echo "Finished building installer"

install_daemon: build build_installer 
	bin/installer --binary-file bin/promgateway --config-file bin/promgateway.conf.json

run_from_build: build
	bin/promgateway --config-file bin/promgateway.conf.json

## Common tasks
doc:
    godoc --http :8080

test:
	go test

new_version:
	./scripts/update_version.sh

config_git_hooks: 
	git config core.hooksPath .githooks
