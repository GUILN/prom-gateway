# Readme

This project aims to provide a `gRPC` server that exposes `prometheus` metrics as `remote procedure calls`
and exposes metrics `scrape` endpoint to `prometheus`. Making prometheus to be used as a `push` instead of `pull` model.

### Running the server:
    - Get the binary from your os. (inside bin folder)
    - Run as specified below (where `--config-file` should be a json file with the same template as indicated by the template in this dir)

```bash
promexporter --metrics-handler-address 0.0.0.0 \
		--metrics-handler-port 50051 \
		--prometheus-metrics-address 0.0.0.0 \
		--prometheus-metrics-port 8080 \
```

**OR**


```bash
promexporter --config-file promgateway.conf.json
```
### Using the server:

- To test
    - You can use cli tools to connect to `gRPC` server and describe the services, such as [evans](https://github.com/ktr0731/evans).
    - After invoking some metrics service, such as `IncrementCounter` you should be able to access `<host>:<port>/metrics` endpoint and vrify that your metric is being incremented.
    - If you prefer you can start the server and run test assets:
        - go to [test_assets folder](./test_assets), 
        - run `docker-compose up`, to run prometheus (make sure the address you started the server is the same as configured in prometheus config)
        - go to [test_assets/bin folder](./test_assets/bin) and run the test client, passing as the parameter the `gRPC` address configured when running the server.

- To use in your application:
    - If you use `golang` the `client code` is already generated for you to use in [proto folder](./proto)
    - If you use other language, you can generate the `client code`:
        - Install `protoc` protocol buffer compiler by following [this article](https://grpc.io/docs/protoc-installation/)
        - Compile client code. You can follow this [manual](https://helpmanual.io/man1/protoc/)
