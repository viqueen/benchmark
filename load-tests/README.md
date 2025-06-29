# k6 Load Tests with Grafana

Load testing setup for the Content API using k6 and Grafana visualization.

## Quick Start

1. **Start monitoring stack:**
   ```bash
   docker-compose up influxdb grafana -d
   ```

2. **Run load tests:**
   ```bash
   docker-compose --profile test up k6
   ```

3. **View results in Grafana:**
   - Open http://localhost:3000
   - Login: admin/admin
   - Navigate to the k6 Load Testing Dashboard

## Available Commands

- **Smoke test (1 user, 30s):**
  ```bash
  npm run test:smoke
  ```

- **Load test (10 users, 5m):**
  ```bash
  npm run test:load
  ```

- **Stress test (50 users, 10m):**
  ```bash
  npm run test:stress
  ```

## Clean Environment

To reset all data and start fresh:
```bash
docker-compose down
rm -rf .local/
docker-compose up influxdb grafana -d
```

## Architecture

- **InfluxDB**: Stores k6 metrics data
- **Grafana**: Visualizes metrics with pre-configured dashboard
- **k6**: Runs load tests against the Content API
- **API**: Go-based Content Service (from packages/go-basic)

All monitoring data is stored in `.local/` directory for easy cleanup.