package controllers

import (
	"ExpensesTracker/internal/api/http/services"
	"ExpensesTracker/pkg/api/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	services.DeleteExpensesSummary(uint(id))
}

func GetExpensesSummary(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	services.GetExpensesSummary(uint(id))
}
