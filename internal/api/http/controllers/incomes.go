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
func UpsertIncome(c *gin.Context) {
	summaryExpenseId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadPathParameterError("Income"))
	}
	SavingId := c.Param("resourceStringId")
	var updateIncome requests.CreateIncomeRequest
	err = c.BindJSON(&updateIncome)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadRequestBodyError("Income"))
	}
	res := services.UpsertIncome(uint(summaryExpenseId), SavingId, updateIncome)
	c.JSON(http.StatusOK, res)
}

func DeleteIncome(c *gin.Context) {
	expenseSummaryId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadPathParameterError("Income"))
	}
	expenseId := c.Param("resourceStringId")
	err = services.DeleteIncome(uint(expenseSummaryId), expenseId)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewNotFoundError("Income"))
	}
	c.JSON(http.StatusNoContent, nil)
}
