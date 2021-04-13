package exporter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"

	"github.com/danopstech/starlink_exporter/pkg/spacex.com/api/device"
)

const (
	dishAddress = "192.168.100.1:9200"
	namespace   = "starlink"
)

var (
	dishUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "up"),
		"Was the last query of Starlink successful.",
		nil, nil,
	)
	dishScrapeDurationSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "scrape_duration_seconds"),
		"Time to scrape metrics from starlink dish",
		nil, nil,
	)
	dishDeviceInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "dishDeviceInfo"),
		"Running software versions and IDs of hardware",
		[]string{"device_id", "hardware_version", "software_version", "country_code", "utc_offset"}, nil,
	)
	dishState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "state"),
		"The Current dishState of the Dish (Unknown, Booting, Searching, Connected).",
		nil, nil,
	)
	dishUptimeSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "uptime_seconds"),
		"Dish running time",
		nil, nil,
	)
	dishPopPingDropRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "pop_ping_drop_ratio"),
		"Percent of pings dropped",
		nil, nil,
	)
	dishPopPingLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "pop_ping_latency_seconds"),
		"Latency of connection in seconds",
		nil, nil,
	)
	dishSnr = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "snr"),
		"Signal strength of the connection",
		nil, nil,
	)
	dishUplinkThroughputBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "uplink_throughput_bytes"),
		"Amount of bandwidth in bytes per second upload",
		nil, nil,
	)
	dishDownlinkThroughputBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "downlink_throughput_bytes"),
		"Amount of bandwidth in bytes per second download",
		nil, nil,
	)
	dishCurrentlyObstructed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "currently_obstructed"),
		"Status of view of the sky",
		nil, nil,
	)
	dishFractionObstructionRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "fraction_obstruction_ratio"),
		"Percentage of obstruction",
		nil, nil,
	)
	dishLast24hObstructedSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "last_24h_obstructed_seconds"),
		"Number of seconds view of sky has been obstructed in the last 24hours",
		nil, nil,
	)
	dishValidSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "valid_seconds"),
		"Unknown",
		nil, nil,
	)
	dishSecondsToFirstNonemptySlot = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "first_nonempty_slot_seconds"),
		"Seconds to next non empty slot",
		nil, nil,
	)
	dishWedgeFractionObstructionRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "wedge_fraction_obstruction_ratio"),
		"Percentage of obstruction per wedge section",
		[]string{"wedge"}, nil,
	)
	dishWedgeAbsFractionObstructionRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "wedge_abs_fraction_obstruction_ratio"),
		"Percentage of Absolute fraction per wedge section",
		[]string{"wedge"}, nil,
	)
	dishAlertMotorsStuck = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_motors_stuck"),
		"Status of motor stuck",
		nil, nil,
	)
	dishAlertThermalThrottle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_thermal_throttle"),
		"Status of thermal throttling",
		nil, nil,
	)
	dishAlertThermalShutdown = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_thermal_shutdown"),
		"Status of thermal shutdown",
		nil, nil,
	)
	dishAlertMastNotNearVertical = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_mast_not_near_vertical"),
		"Status of mast position",
		nil, nil,
	)
	dishUnexpectedLocation = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_unexpected_location"),
		"Status of location",
		nil, nil,
	)
	dishSlowEthernetSpeeds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_slow_eth_speeds"),
		"Status of ethernet",
		nil, nil,
	)
)

// Exporter collects Starlink stats from the Dish and exports them using
// the prometheus metrics package.
type Exporter struct {
	Conn   *grpc.ClientConn
	Client device.DeviceClient
}

// New returns an initialized Exporter.
func New() (*Exporter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	conn, err := grpc.DialContext(ctx, dishAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("error creating underlying gRPC connection to starlink dish: %s", err.Error())
	}

	return &Exporter{
		Conn:   conn,
		Client: device.NewDeviceClient(conn),
	}, nil
}

// Describe describes all the metrics ever exported by the Starlink exporter. It
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- dishUp
	ch <- dishScrapeDurationSeconds
	ch <- dishDeviceInfo
	ch <- dishState
	ch <- dishUptimeSeconds
	ch <- dishPopPingDropRatio
	ch <- dishPopPingLatencySeconds
	ch <- dishSnr
	ch <- dishUplinkThroughputBytes
	ch <- dishDownlinkThroughputBytes
	ch <- dishCurrentlyObstructed
	ch <- dishFractionObstructionRatio
	ch <- dishLast24hObstructedSeconds
	ch <- dishValidSeconds
	ch <- dishSecondsToFirstNonemptySlot
	ch <- dishWedgeFractionObstructionRatio
	ch <- dishWedgeAbsFractionObstructionRatio
	ch <- dishAlertMotorsStuck
	ch <- dishAlertThermalThrottle
	ch <- dishAlertThermalShutdown
	ch <- dishAlertMastNotNearVertical
	ch <- dishUnexpectedLocation
	ch <- dishSlowEthernetSpeeds
}

// Collect fetches the stats from Starlink dish and delivers them
// as Prometheus metrics. It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()
	ok := e.collectDish(ch)
	dishDuration := time.Since(start).Seconds()

	if ok {
		ch <- prometheus.MustNewConstMetric(
			dishUp, prometheus.GaugeValue, 1.0,
		)
		ch <- prometheus.MustNewConstMetric(
			dishScrapeDurationSeconds, prometheus.GaugeValue, dishDuration,
		)
	} else {
		ch <- prometheus.MustNewConstMetric(
			dishUp, prometheus.GaugeValue, 0.0,
		)
	}
}

func (e *Exporter) collectDish(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetStatus{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("failed to collect metrics from dish: %s", err.Error())
		return false
	}

	dishStatus := resp.GetDishGetStatus()
	deviceInfo := dishStatus.GetDeviceInfo()
	obstructions := dishStatus.GetObstructionStats()
	alerts := dishStatus.GetAlerts()

	ch <- prometheus.MustNewConstMetric(
		dishDeviceInfo, prometheus.GaugeValue, 1.00,
		deviceInfo.GetId(),
		deviceInfo.GetHardwareVersion(),
		deviceInfo.GetSoftwareVersion(),
		deviceInfo.GetCountryCode(),
		fmt.Sprint(deviceInfo.GetUtcOffsetS()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishState, prometheus.GaugeValue, float64(dishStatus.GetState().Number()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishUptimeSeconds, prometheus.GaugeValue, float64(dishStatus.GetDeviceState().GetUptimeS()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishSecondsToFirstNonemptySlot, prometheus.GaugeValue, float64(dishStatus.GetSecondsToFirstNonemptySlot()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishPopPingDropRatio, prometheus.GaugeValue, float64(dishStatus.GetPopPingDropRate()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishPopPingLatencySeconds, prometheus.GaugeValue, float64(dishStatus.GetPopPingLatencyMs()/1000),
	)

	ch <- prometheus.MustNewConstMetric(
		dishSnr, prometheus.GaugeValue, float64(dishStatus.GetSnr()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishUplinkThroughputBytes, prometheus.GaugeValue, float64(dishStatus.GetUplinkThroughputBps()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishDownlinkThroughputBytes, prometheus.GaugeValue, float64(dishStatus.GetDownlinkThroughputBps()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishCurrentlyObstructed, prometheus.GaugeValue, flool(obstructions.GetCurrentlyObstructed()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishFractionObstructionRatio, prometheus.GaugeValue, float64(obstructions.GetFractionObstructed()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishLast24hObstructedSeconds, prometheus.GaugeValue, float64(obstructions.GetLast_24HObstructedS()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishValidSeconds, prometheus.GaugeValue, float64(obstructions.GetValidS()),
	)

	for i, v := range obstructions.GetWedgeFractionObstructed() {
		ch <- prometheus.MustNewConstMetric(
			dishWedgeFractionObstructionRatio, prometheus.GaugeValue, float64(v),
			strconv.Itoa(i),
		)
	}

	for i, v := range obstructions.GetWedgeAbsFractionObstructed() {
		ch <- prometheus.MustNewConstMetric(
			dishWedgeAbsFractionObstructionRatio, prometheus.GaugeValue, float64(v),
			strconv.Itoa(i),
		)
	}

	ch <- prometheus.MustNewConstMetric(
		dishAlertMotorsStuck, prometheus.GaugeValue, flool(alerts.GetMotorsStuck()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishAlertThermalThrottle, prometheus.GaugeValue, flool(alerts.GetThermalThrottle()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishAlertThermalShutdown, prometheus.GaugeValue, flool(alerts.GetThermalShutdown()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishAlertMastNotNearVertical, prometheus.GaugeValue, flool(alerts.GetMastNotNearVertical()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishUnexpectedLocation, prometheus.GaugeValue, flool(alerts.GetUnexpectedLocation()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishSlowEthernetSpeeds, prometheus.GaugeValue, flool(alerts.GetSlowEthernetSpeeds()),
	)

	return true
}

func flool(b bool) float64 {
	if b {
		return 1.00
	}
	return 0.00
}
