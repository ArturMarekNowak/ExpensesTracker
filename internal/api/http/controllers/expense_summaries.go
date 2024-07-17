package controllers

import (
	"ExpensesTracker/internal/api/http/services"
	"ExpensesTracker/pkg/api/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateExpensesSummary(c *gin.Context) {
	var createExpenseSummaryRequest requests.CreateExpenseSummaryRequest
	err := c.BindJSON(&createExpenseSummaryRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	expensesSummary := services.CreateExpensesSummary(createExpenseSummaryRequest)
	c.JSON(http.StatusCreated, expensesSummary)
}

func DeleteExpensesSummary(c *gin.Context) {

}

func GetExpensesSummary(c *gin.Context) {

}
