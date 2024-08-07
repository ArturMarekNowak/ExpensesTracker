package services

import (
	"ExpensesTracker/pkg/api/http/requests"
	"ExpensesTracker/pkg/api/http/responses"
	"ExpensesTracker/pkg/database"
	"ExpensesTracker/pkg/database/entities"
	"errors"
)

// Source: https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
func UpsertIncome(summaryExpenseId uint, name string, updateSaving requests.CreateIncomeRequest) responses.IncomeResponse {
	income := entities.Income{
		Name:              name,
		ExpensesSummaryId: summaryExpenseId,
		Value:             updateSaving.Value,
	}
	db := database.OpenConnection()
	if db.Model(&income).Where("name = ?", income.Name).Updates(&income).RowsAffected == 0 {
		db.Create(&income)
	}
	return responses.IncomeResponse{
		Name:  income.Name,
		Value: income.Value,
	}
}

func DeleteIncome(expensesSummaryId uint, incomeId string) error {
	db := database.OpenConnection()
	res := db.Delete(entities.Income{
		ExpensesSummaryId: expensesSummaryId,
		Name:              incomeId,
	})
	if res.RowsAffected == 0 {
		return errors.New("No data found")
	}
	return nil
}
