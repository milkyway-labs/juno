package logging

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// StartHeight represents the Telemetry counter used to set the start height of the parsing
var StartHeight = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_initial_height",
		Help: "Initial parsing height.",
	},
)

// WorkerCount represents the Telemetry counter used to track the worker count
var WorkerCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_worker_count",
		Help: "Number of active workers.",
	},
)

// WorkerHeight represents the Telemetry counter used to track the last indexed height for each worker
var WorkerHeight = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "juno_last_indexed_height",
		Help: "Height of the last indexed block.",
	},
	[]string{"worker_index", "chain_id"},
)

// ErrorCount represents the Telemetry counter used to track the number of errors emitted
var ErrorCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_error_count",
		Help: "Total number of errors emitted.",
	},
)

var RPCRequestErrors = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_rpc_errors_total",
		Help: "Total number of errors occurred during RPC requests",
	},
)

var DBOperationErrors = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_db_errors_total",
		Help: "Total number of errors occurred during database operations",
	},
)

var ProcessBlockErrorCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "juno_block_errors_total",
		Help: "Total number of errors per block",
	},
	[]string{"block"},
)

// DBLatestHeight represents the Telemetry counter used to track the last indexed height in the database
var DBLatestHeight = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "juno_db_latest_height",
		Help: "Latest block height in the database.",
	},
	[]string{"db_latest_height"},
)

// SignalRPCRequestError signal that a new rpc request error occurred
func SignalRPCRequestError() {
	RPCRequestErrors.Inc()
}

// SignalDBOperationError signal that a new error occurred while interacting
// with the database
func SignalDBOperationError() {
	DBOperationErrors.Inc()
}

// SignalBlockError increments the error counter for the given block
func SignalBlockError(blockHeight int64) {
	blockStr := fmt.Sprintf("%d", blockHeight)
	ProcessBlockErrorCount.WithLabelValues(blockStr).Inc()
	prometheus.MustRegister()
}

func init() {
	prometheus.MustRegister(StartHeight)
	prometheus.MustRegister(WorkerCount)
	prometheus.MustRegister(WorkerHeight)
	prometheus.MustRegister(ErrorCount)
	prometheus.MustRegister(DBLatestHeight)
	prometheus.MustRegister(RPCRequestErrors)
	prometheus.MustRegister(DBOperationErrors)
	prometheus.MustRegister(ProcessBlockErrorCount)
}
