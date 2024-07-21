package endpoints

import (
	"ExpensesTracker/internal/api/http/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureEndpoints() {
	router := gin.Default()
	router.POST("/expensesSummaries", controllers.CreateExpensesSummary)
	router.GET("/expensesSummaries/:resourceIntegerId", controllers.GetExpensesSummary)
	router.DELETE("/expenseSummaries/:resourceIntegerId", controllers.DeleteExpensesSummary)
	router.PUT("/expensesSummaries/:resourceIntegerId/expenses/:resourceStringId", controllers.UpdateExpense)
	router.DELETE("/expensesSummaries/:resourceIntegerId/expenses/:resourceStringId", controllers.DeleteExpense)
	router.PUT("/expensesSummaries/:resourceIntegerId/incomes/:resourceStringId", controllers.UpdateIncome)
	router.DELETE("/expensesSummaries/:resourceIntegerId/incomes/:resourceStringId", controllers.DeleteIncome)
	router.PUT("/expensesSummaries/:resourceIntegerId/savings/:resourceStringId", controllers.UpdateSaving)
	router.DELETE("/expensesSummaries/:resourceIntegerId/savings/:resourceStringId", controllers.DeleteSaving)
	err := router.Run(":8080")
	if err != nil {
		panic("Failed to start host")
	}
}
