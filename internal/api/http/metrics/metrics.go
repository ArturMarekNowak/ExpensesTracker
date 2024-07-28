package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func ConfigureMetrics() {

	http.Handle("/metrics", promhttp.Handler())
}
