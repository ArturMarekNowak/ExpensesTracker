package endpoints

import (
	"ExpensesTracker/internal/api/http/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureEndpoints() {
	router := gin.Default()
	router.POST("/expenseSummaries", controllers.CreateExpensesSummary)
	router.GET("/expenseSummaries/:resourceIntegerId", controllers.GetExpensesSummary)
	router.DELETE("/expenseSummaries/:resourceIntegerId", controllers.DeleteExpensesSummary)
	router.PUT("/expensesSummary/:resourceIntegerId/expenses/:resourceStringId", controllers.UpdateExpense)
	router.DELETE("/expensesSummary/:resourceIntegerId/expenses/:resourceStringId", controllers.DeleteExpense)
	router.PUT("/expensesSummary/:resourceIntegerId/incomes/:resourceStringId", controllers.UpdateIncome)
	router.DELETE("/expensesSummary/:resourceIntegerId/incomes/:resourceStringId", controllers.DeleteIncome)
	router.PUT("/expensesSummary/:resourceIntegerId/savings/:resourceStringId", controllers.UpdateSaving)
	router.DELETE("/expensesSummary/:resourceIntegerId/savings/:resourceStringId", controllers.DeleteSaving)
	err := router.Run(":8080")
	if err != nil {
		panic("Failed to start host")
	}
}
