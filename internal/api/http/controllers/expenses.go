package controllers

import (
	"ExpensesTracker/internal/api/http/services"
	"ExpensesTracker/pkg/api/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Source: https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
func UpdateExpense(c *gin.Context) {
	summaryExpenseId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	SavingId := c.Param("resourceStringId")
	var updateIncome requests.CreateExpenseRequest
	err = c.BindJSON(&updateIncome)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	res := services.UpdateExpense(uint(summaryExpenseId), SavingId, updateIncome)
	c.JSON(http.StatusOK, res)
}

func DeleteExpense(c *gin.Context) {
	expenseSummaryId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	expenseId := c.Param("resourceStringId")
	err = services.DeleteExpense(uint(expenseSummaryId), expenseId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid path parameter"})
	}
	c.JSON(http.StatusNoContent, nil)
}
