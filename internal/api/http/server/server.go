package server

import (
	"ExpensesTracker/pkg/database"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func Start() {

	database.RunMigrations()

	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("failed to start host")
	}
}
