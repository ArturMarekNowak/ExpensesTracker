package controllers

import (
	"ExpensesTracker/internal/api/http/services"
	"ExpensesTracker/pkg/api/http/requests"
	"ExpensesTracker/pkg/api/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateExpensesSummary(c *gin.Context) {
	var createExpenseSummaryRequest requests.CreateExpenseSummaryRequest
	err := c.BindJSON(&createExpenseSummaryRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadRequestBodyError("ExpensesSummary"))
	}
	expensesSummary := services.CreateExpensesSummary(createExpenseSummaryRequest)
	c.JSON(http.StatusCreated, expensesSummary)
}

func DeleteExpensesSummary(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadPathParameterError("ExpensesSummary"))
	}
	err = services.DeleteExpensesSummary(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewNotFoundError("ExpensesSummary"))
	}
	c.JSON(http.StatusNoContent, nil)
}

func GetExpensesSummary(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadPathParameterError("ExpensesSummary"))
	}
	res, err := services.GetExpensesSummary(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewNotFoundError("ExpensesSummary"))
	}
	c.JSON(http.StatusOK, res)
}
