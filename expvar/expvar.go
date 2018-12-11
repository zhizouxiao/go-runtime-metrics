package expvar

import (
	"expvar"
	"os"

	"github.com/zhizouxiao/go-runtime-metrics/influxdb"
)

const defaultMeasurement = "go_runtime_metrics"

func init() {
	expvar.Publish(os.Args[0], influxdb.Metrics(defaultMeasurement))
}
