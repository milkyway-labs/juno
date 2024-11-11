package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

// StartHeight represents the Telemetry counter used to set the start height of the parsing
var StartHeight = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_start_height",
		Help: "Initial parsing height.",
	},
)

// WorkersCount represents the Telemetry counter used to track the number of active workers
var WorkersCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_active_workers",
		Help: "Number of active workers.",
	},
)

// LatestIndexedHeightByWorker represents the Telemetry counter used to track the last indexed height for each worker
var LatestIndexedHeightByWorker = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "juno_latest_indexed_height",
		Help: "Height of the last indexed block.",
	},
	[]string{"worker_index"},
)

func init() {
	prometheus.MustRegister(StartHeight)
	prometheus.MustRegister(WorkersCount)
	prometheus.MustRegister(LatestIndexedHeightByWorker)
}
