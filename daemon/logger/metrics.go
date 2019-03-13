package logger // import "github.com/docker/docker/daemon/logger"

import (
	"github.com/docker/go-metrics"
)

var (
	logWritesFailedCount           metrics.Counter
	logReadsFailedCount            metrics.Counter
	totalPartialLogs               metrics.Counter
	ringBufferEnqueueFailedCount   metrics.Counter
	ringBufferLogWritesFailedCount metrics.Counter
)

func init() {
	loggerMetrics := metrics.NewNamespace("logger", "", nil)

	// NOTE: For our case, this should just be errors when the ringlogger queue is closed
	logWritesFailedCount = loggerMetrics.NewCounter("log_write_operations_failed", "Number of log write operations that failed")
	logReadsFailedCount = loggerMetrics.NewCounter("log_read_operations_failed", "Number of log reads from container stdio that failed")
	totalPartialLogs = loggerMetrics.NewCounter("log_entries_size_greater_than_buffer", "Number of log entries which are larger than the log buffer")

	ringBufferEnqueueFailedCount = loggerMetrics.NewCounter("ring_buffer_enqueue_operations_failed", "Number of log entries dropped because they were larger than the available ring buffer")
	// NOTE: AFAICT the only error we would see here is if there is a failure encoding the log into protobuf in the plugin adapter
	ringBufferLogWritesFailedCount = loggerMetrics.NewCounter("ring_buffer_log_write_operations_failed", "Number of log writes to the underlying log driver that failed")

	metrics.Register(loggerMetrics)
}
