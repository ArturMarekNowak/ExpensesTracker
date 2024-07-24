package controllers

import (
	"ExpensesTracker/internal/api/http/services"
	"ExpensesTracker/pkg/api/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Source: https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
func UpdateIncome(c *gin.Context) {
	summaryExpenseId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	SavingId := c.Param("resourceStringId")
	var updateIncome requests.CreateIncomeRequest
	err = c.BindJSON(&updateIncome)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	res := services.UpdateIncome(uint(summaryExpenseId), SavingId, updateIncome)
	c.JSON(http.StatusOK, res)
}

func DeleteIncome(c *gin.Context) {
	expenseSummaryId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	expenseId := c.Param("resourceStringId")
	err = services.DeleteIncome(uint(expenseSummaryId), expenseId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid path parameter"})
	}
	c.JSON(http.StatusNoContent, nil)
}
