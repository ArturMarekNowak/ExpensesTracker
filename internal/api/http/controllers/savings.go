package controllers

import (
	"ExpensesTracker/internal/api/http/services"
	"ExpensesTracker/pkg/api/http/requests"
	"ExpensesTracker/pkg/api/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateSaving(c *gin.Context) {
	summaryExpenseId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadPathParameterError("Saving"))
	}
	SavingId := c.Param("resourceStringId")
	var updateSaving requests.CreateSavingRequest
	err = c.BindJSON(&updateSaving)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadRequestBodyError("Saving"))
	}
	res := services.UpdateSaving(uint(summaryExpenseId), SavingId, updateSaving)
	c.JSON(http.StatusOK, res)
}

func DeleteSaving(c *gin.Context) {
	expenseSummaryId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadPathParameterError("Saving"))
	}
	expenseId := c.Param("resourceStringId")
	err = services.DeleteSaving(uint(expenseSummaryId), expenseId)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewNotFoundError("Saving"))
	}
	c.JSON(http.StatusNoContent, nil)
}
