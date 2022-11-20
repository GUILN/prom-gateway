# Prometheus exporter

Despite of being a prometheus anti-pattern, exporting metrics somethimes is simpler to get some metrics up and running,
specially if you are working in an environment where you have freedom to quickly implemet ephemeral systems like batching and serverless very quickly.


This daemon captures receives metrics from services and exposes them for prometheus in a metrics `/metrics` endpoint to be scraped by prometheus `pull system` server.
