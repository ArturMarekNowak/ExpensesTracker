package controllers

import (
	"ExpensesTracker/internal/api/http/services"
	"ExpensesTracker/pkg/api/http/requests"
	"ExpensesTracker/pkg/api/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Source: https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
func UpsertExpense(c *gin.Context) {
	summaryExpenseId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadPathParameterError("Expense"))
	}
	SavingId := c.Param("resourceStringId")
	var updateIncome requests.CreateExpenseRequest
	err = c.BindJSON(&updateIncome)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadRequestBodyError("Expense"))
	}
	res := services.UpsertExpense(uint(summaryExpenseId), SavingId, updateIncome)
	c.JSON(http.StatusOK, res)
}

func DeleteExpense(c *gin.Context) {
	expenseSummaryId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadPathParameterError("Expense"))
	}
	expenseId := c.Param("resourceStringId")
	err = services.DeleteExpense(uint(expenseSummaryId), expenseId)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewNotFoundError("Expense"))
	}
	c.JSON(http.StatusNoContent, nil)
}
