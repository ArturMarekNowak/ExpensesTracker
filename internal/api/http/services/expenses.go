package services

import (
	"ExpensesTracker/pkg/api/http/requests"
	"ExpensesTracker/pkg/api/http/responses"
	"ExpensesTracker/pkg/database"
	"ExpensesTracker/pkg/database/entities"
	"errors"
)

// Source: https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
func UpdateExpense(summaryExpenseId uint, name string, updateSaving requests.CreateExpenseRequest) responses.ExpenseResponse {
	expense := entities.Expense{
		Name:              name,
		ExpensesSummaryId: summaryExpenseId,
		Value:             updateSaving.Value,
	}
	db := database.OpenConnection()
	if db.Model(&expense).Where("name = ?", expense.Name).Updates(&expense).RowsAffected == 0 {
		db.Create(&expense)
	}
	return responses.ExpenseResponse{
		Name:  expense.Name,
		Value: expense.Value,
	}
}

func DeleteExpense(id string) error {
	db := database.OpenConnection()
	res := db.Delete(entities.Expense{Name: id})
	if res.RowsAffected == 0 {
		return errors.New("No data found")
	}
	return nil
}
