# Prometheus exporter

Despite of being a prometheus anti-pattern, exporting metrics somethimes is simpler to get some metrics up and running,
specially if you are working in an environment where you have freedom to quickly implemet ephemeral systems like batching and serverless very quickly.


This daemon captures receives metrics from services and exposes them for prometheus in a metrics `/metrics` endpoint to be scraped by prometheus `pull system` server.

## !IMPORTANT: For contributors
If you are a contributor, please follow the steps below to enable `git hooks` used by this project:

- Run `make config_git_hooks` to set the `git hooks` folder to [project's git hook folder](./.githooks). 

### Testing locally

Run prometheus:
```bash
docker-compose up
```

Run daemon process:
```bash
make run_daemon_process
```

Use `grpcurl`:
```bash
grpcurl -plaintext -d '{"counter_name": "my_app_some_counter", "counter_help": "some counter"}' localhost:50051 Metrics/IncrementCounter
```
