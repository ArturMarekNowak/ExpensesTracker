package endpoints

import (
	"ExpensesTracker/internal/api/http/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureEndpoints() {
	router := gin.Default()
	router.POST("/expensesSummaries", controllers.CreateExpensesSummary)
	router.GET("/expensesSummaries/:resourceIntegerId", controllers.GetExpensesSummary)
	router.DELETE("/expensesSummaries/:resourceIntegerId", controllers.DeleteExpensesSummary)
	router.PUT("/expensesSummaries/:resourceIntegerId/expenses/:resourceStringId", controllers.UpsertExpense)
	router.DELETE("/expensesSummaries/:resourceIntegerId/expenses/:resourceStringId", controllers.DeleteExpense)
	router.PUT("/expensesSummaries/:resourceIntegerId/incomes/:resourceStringId", controllers.UpsertIncome)
	router.DELETE("/expensesSummaries/:resourceIntegerId/incomes/:resourceStringId", controllers.DeleteIncome)
	router.PUT("/expensesSummaries/:resourceIntegerId/savings/:resourceStringId", controllers.UpsertSaving)
	router.DELETE("/expensesSummaries/:resourceIntegerId/savings/:resourceStringId", controllers.DeleteSaving)
	err := router.Run(":8080")
	if err != nil {
		panic("Failed to start host")
	}
}
