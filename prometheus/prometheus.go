package prometheus

import (
	"fmt"

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

// ErrorsCount represents the Telemetry counter used to track the number of errors emitted
var ErrorsCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_errors",
		Help: "Total number of errors emitted.",
	},
)

// RPCRequestErrorsCount represents the Telemetry counter used to track the number of RPC request errors
var RPCRequestErrorsCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_rpc_errors",
		Help: "Total number of errors occurred during RPC requests",
	},
)

// DBOperationErrorsCount represents the Telemetry counter used to track the number of database operation errors
var DBOperationErrorsCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_db_errors",
		Help: "Total number of errors occurred during database operations",
	},
)

// BlockErrorsCount represents the Telemetry counter used to track the number of errors per block
var BlockErrorsCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "juno_block_errors",
		Help: "Total number of errors per block",
	},
	[]string{"block"},
)

// SignalRPCRequestError signal that a new rpc request error occurred
func SignalRPCRequestError() {
	RPCRequestErrorsCount.Inc()
}

// SignalDBOperationError signal that a new error occurred while interacting
// with the database
func SignalDBOperationError() {
	DBOperationErrorsCount.Inc()
}

// SignalBlockError increments the error counter for the given block
func SignalBlockError(blockHeight int64) {
	blockStr := fmt.Sprintf("%d", blockHeight)
	BlockErrorsCount.WithLabelValues(blockStr).Inc()
	prometheus.MustRegister()
}

func init() {
	prometheus.MustRegister(StartHeight)
	prometheus.MustRegister(WorkersCount)
	prometheus.MustRegister(LatestIndexedHeightByWorker)
	prometheus.MustRegister(ErrorsCount)
	prometheus.MustRegister(RPCRequestErrorsCount)
	prometheus.MustRegister(DBOperationErrorsCount)
	prometheus.MustRegister(BlockErrorsCount)
}
