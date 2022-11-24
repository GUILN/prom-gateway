FROM golang:1.19-alpine as builder 

RUN mkdir /app
COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o promgateway ./daemon/main.go

RUN chmod +x /app/promgateway

# Run tests from docker

## Set test variables

ENV HANDLER_PORT=50051
ENV HANDLER_ADDRESS=0.0.0.0
ENV METRICS_PORT=8080
ENV METRICS_ADDRESS=0.0.0.0

## Running tests
### Command below starst the application with nohup and runs the integration tests against it.
RUN nohup /app/promgateway --metrics-handler-address ${HANDLER_ADDRESS} --metrics-handler-port ${HANDLER_PORT} --prometheus-metrics-address ${METRICS_ADDRESS} --prometheus-metrics-port ${METRICS_PORT} & CGO_ENABLED=0 go test ./test/

# Generates a tiny docker image.
FROM alpine:latest

RUN mkdir /app
COPY --from=builder /app/promgateway /app
CMD ["/app/promgateway", "--metrics-handler-address", $MADDR, "--metrics-handler-port", $MPORT, "--prometheus-metrics-address", $PADDR, "--prometheus-metrics-port", $PPORT]
