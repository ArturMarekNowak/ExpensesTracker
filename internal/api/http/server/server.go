package server

import (
	"ExpensesTracker/internal/api/http/endpoints"
	"ExpensesTracker/internal/api/http/metrics"
	"ExpensesTracker/pkg/database"
)

func Start() {

	database.RunMigrations()
	metrics.ConfigureMetrics()
	endpoints.ConfigureEndpoints()
}
