# 04 Logs and Traces

Logs and Traces can also be templated. This is built ontop of the previous examples.

`docker run -v ${PWD}/:/etc/grafana grafana/agentctl:latest template-parse file:///etc/grafana/04_config.yml`

## Dynamic Configuration

[config.yml](04_config.yml)

Tells the Grafana Agent where to load files from.

## Logs

Logs are loaded from a template matching `logs-*.yml`. There can ONLY be 1 template loaded

[logs-1.yml](04_assets/logs-1.yml)

```yaml
configs:
  - name: test_logs
    positions:
      filename: /tmp/positions.yaml
    scrape_configs:
      - job_name: test
        pipeline_stages:
          - regex:
            source: filename
            expression: '\\temp\\Logs\\(?P<log_app>.+?)\\'
```

[traces.yml](04_assets/traces-1.yml)

```yaml
configs:
  - name: test_traces
    automatic_logging:
      backend: stdout
      loki_name: default
      spans: true
```

## Final

[final.yml](04_assets/final.yml)

