package server

import (
	"ExpensesTracker/internal/api/http/endpoints"
	"ExpensesTracker/internal/api/http/metrics"
	"ExpensesTracker/internal/helpers"
	"ExpensesTracker/pkg/database"
)

func Start() {

	helpers.LoadEnvironmentalVariables()
	database.RunMigrations()
	metrics.ConfigureMetrics()
	endpoints.ConfigureEndpoints()
}
