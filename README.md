<p align="center">
  <img alt="logo" src="https://github.com/danopstech/starlink_exporter/raw/main/.docs/assets/logo.jpg" height="150" />
  <h3 align="center">Starlink Prometheus Exporter</h3>
</p>

---
A [Starlink](https://www.starlink.com/) exporter for Prometheus. Not affiliated with or acting on behalf of Starlink(™)

[![goreleaser](https://github.com/danopstech/starlink_exporter/actions/workflows/release.yaml/badge.svg)](https://github.com/danopstech/starlink_exporter/actions/workflows/release.yaml)
[![build](https://github.com/danopstech/starlink_exporter/actions/workflows/build.yaml/badge.svg)](https://github.com/danopstech/starlink_exporter/actions/workflows/build.yaml)
[![License](https://img.shields.io/github/license/danopstech/starlink_exporter)](/LICENSE)
[![Release](https://img.shields.io/github/release/danopstech/starlink_exporter.svg)](https://github.com/danopstech/starlink_exporter/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/danopstech/starlink_exporter)
![os/arch](https://img.shields.io/badge/os%2Farch-amd64-yellow)
![os/arch](https://img.shields.io/badge/os%2Farch-arm64-yellow)
![os/arch](https://img.shields.io/badge/os%2Farch-armv7-yellow)
[![Go Report Card](https://goreportcard.com/badge/github.com/danopstech/starlink_exporter)](https://goreportcard.com/report/github.com/danopstech/starlink_exporter)

If you would like a pre-packaged system to monitor you Starlink system please check out https://github.com/danopstech/starlink. It includes this exporter, speedtest_exporter, blackbox_exporter, Grafana and Prometheus in one Docker Compose file.

## Usage:

### Flags

`starlink_exporter` is configured through the use of optional command line flags

```bash
$ ./starlink_exporter --help
Usage of starlink_exporter
  -address string
        IP address and port to reach dish (default "192.168.100.1:9200")
  -port string
        listening port to expose metrics on (default "9817")

```

### Binaries

For pre-built binaries please take a look at the [releases](https://github.com/danopstech/starlink_exporter/releases).

```bash
./starlink_exporter [flags]
```

### Docker

Docker Images can be found at [GitHub Container Registry](https://github.com/orgs/danopstech/packages/container/package/starlink_exporter) & [Dockerhub](https://hub.docker.com/r/danopstech/starlink_exporter).

Example:
```bash
docker pull ghcr.io/danopstech/starlink_exporter:latest

docker run \
  -p 9817:9817 \
  ghcr.io/danopstech/starlink_exporter:latest [flags]
```

### Setup Prometheus to scrape `starlink_exporter`

Configure [Prometheus](https://prometheus.io/) to scrape metrics from localhost:9817/metrics

```yaml
...
scrape_configs
    - job_name: starlink
      scrape_interval: 3s
      scrape_timeout:  3s
      static_configs:
        - targets: ['localhost:9817']
...
```

## Exported Metrics:

```text
# HELP starlink_dish_alert_mast_not_near_vertical Status of mast position
# TYPE starlink_dish_alert_mast_not_near_vertical gauge
# HELP starlink_dish_alert_motors_stuck Status of motor stuck
# TYPE starlink_dish_alert_motors_stuck gauge
# HELP starlink_dish_alert_slow_eth_speeds Status of ethernet
# TYPE starlink_dish_alert_slow_eth_speeds gauge
# HELP starlink_dish_alert_thermal_shutdown Status of thermal shutdown
# TYPE starlink_dish_alert_thermal_shutdown gauge
# HELP starlink_dish_alert_thermal_throttle Status of thermal throttling
# TYPE starlink_dish_alert_thermal_throttle gauge
# HELP starlink_dish_alert_unexpected_location Status of location
# TYPE starlink_dish_alert_unexpected_location gauge
# HELP starlink_dish_backup_beam connected to backup beam
# TYPE starlink_dish_backup_beam gauge
# HELP starlink_dish_bore_sight_azimuth_deg azimuth in degrees
# TYPE starlink_dish_bore_sight_azimuth_deg gauge
# HELP starlink_dish_bore_sight_elevation_deg elevation in degrees
# TYPE starlink_dish_bore_sight_elevation_deg gauge
# HELP starlink_dish_cell_id Cell ID dish is located in
# TYPE starlink_dish_cell_id gauge
# HELP starlink_dish_currently_obstructed Status of view of the sky
# TYPE starlink_dish_currently_obstructed gauge
# HELP starlink_dish_downlink_throughput_bytes Amount of bandwidth in bytes per second download
# TYPE starlink_dish_downlink_throughput_bytes gauge
# HELP starlink_dish_first_nonempty_slot_seconds Seconds to next non empty slot
# TYPE starlink_dish_first_nonempty_slot_seconds gauge
# HELP starlink_dish_fraction_obstruction_ratio Percentage of obstruction
# TYPE starlink_dish_fraction_obstruction_ratio gauge
# HELP starlink_dish_info Running software versions and IDs of hardware
# TYPE starlink_dish_info gauge
# HELP starlink_dish_initial_gateway_id initial gateway id
# TYPE starlink_dish_initial_gateway_id gauge
# HELP starlink_dish_initial_satellite_id initial satellite id
# TYPE starlink_dish_initial_satellite_id gauge
# HELP starlink_dish_last_24h_obstructed_seconds Number of seconds view of sky has been obstructed in the last 24hours
# TYPE starlink_dish_last_24h_obstructed_seconds gauge
# HELP starlink_dish_pop_ping_drop_ratio Percent of pings dropped
# TYPE starlink_dish_pop_ping_drop_ratio gauge
# HELP starlink_dish_pop_ping_latency_seconds Latency of connection in seconds
# TYPE starlink_dish_pop_ping_latency_seconds gauge
# HELP starlink_dish_pop_rack_id pop rack id
# TYPE starlink_dish_pop_rack_id gauge
# HELP starlink_dish_prolonged_obstruction_duration_seconds Average in seconds of prolonged obstructions
# TYPE starlink_dish_prolonged_obstruction_duration_seconds gauge
# HELP starlink_dish_prolonged_obstruction_interval_seconds Average prolonged obstruction interval in seconds
# TYPE starlink_dish_prolonged_obstruction_interval_seconds gauge
# HELP starlink_dish_scrape_duration_seconds Time to scrape metrics from starlink dish
# TYPE starlink_dish_scrape_duration_seconds gauge
# HELP starlink_dish_snr Signal strength of the connection
# TYPE starlink_dish_snr gauge
# HELP starlink_dish_state The current dishState of the Dish (Unknown, Booting, Searching, Connected).
# TYPE starlink_dish_state gauge
# HELP starlink_dish_time_to_slot_end_seconds Seconds left on current slot
# TYPE starlink_dish_time_to_slot_end_seconds gauge
# HELP starlink_dish_up Was the last query of Starlink dish successful.
# TYPE starlink_dish_up gauge
# HELP starlink_dish_uplink_throughput_bytes Amount of bandwidth in bytes per second upload
# TYPE starlink_dish_uplink_throughput_bytes gauge
# HELP starlink_dish_uptime_seconds Dish running time
# TYPE starlink_dish_uptime_seconds gauge
# HELP starlink_dish_valid_seconds Unknown
# TYPE starlink_dish_valid_seconds gauge
# HELP starlink_dish_wedge_abs_fraction_obstruction_ratio Percentage of Absolute fraction per wedge section
# TYPE starlink_dish_wedge_abs_fraction_obstruction_ratio gauge
# HELP starlink_dish_wedge_fraction_obstruction_ratio Percentage of obstruction per wedge section
# TYPE starlink_dish_wedge_fraction_obstruction_ratio gauge
```

## Example Grafana Dashboard:

https://grafana.com/grafana/dashboards/14337

<p align="center">
	<img src="https://github.com/danopstech/starlink_exporter/raw/main/.docs/assets/screenshot.jpg" width="95%">
</p>
