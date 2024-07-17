package endpoints

import (
	"ExpensesTracker/internal/api/http/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureEndpoints() {
	router := gin.Default()
	err := router.Run(":8080")
	router.POST("/expenseSummaries", controllers.CreateExpensesSummary)
	router.GET("/expenseSummaries/:id", controllers.GetExpensesSummary)
	router.DELETE("/expenseSummaries/:id", controllers.DeleteExpensesSummary)
	router.PUT("/expensesSummary/:id/expenses/:id", controllers.UpdateExpense)
	router.DELETE("/expensesSummary/:id/expenses/:id", controllers.DeleteExpense)
	router.PUT("/expensesSummary/:id/incomes/:id", controllers.UpdateIncome)
	router.DELETE("/expensesSummary/:id/incomes/:id", controllers.DeleteIncome)
	router.PUT("/expensesSummary/:id/savings/:id", controllers.UpdateSaving)
	router.DELETE("/expensesSummary/:id/savings/:id", controllers.DeleteSaving)
	if err != nil {
		panic("Failed to start host")
	}
}
