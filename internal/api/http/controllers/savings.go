package controllers

import (
	"ExpensesTracker/internal/api/http/services"
	"ExpensesTracker/pkg/api/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateSaving(c *gin.Context) {
	summaryExpenseId, err := strconv.ParseUint(c.Param("resourceIntegerId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	SavingId := c.Param("resourceStringId")
	var updateSaving requests.CreateSavingRequest
	err = c.BindJSON(&updateSaving)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	res := services.UpdateSaving(uint(summaryExpenseId), SavingId, updateSaving)
	c.JSON(http.StatusOK, res)
}

func DeleteSaving(c *gin.Context) {
	id := c.Param("resourceStringId")
	err := services.DeleteSaving(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid path parameter"})
	}
	c.JSON(http.StatusNoContent, nil)
}
