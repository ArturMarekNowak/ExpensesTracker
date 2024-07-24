package services

import (
	"ExpensesTracker/pkg/api/http/requests"
	"ExpensesTracker/pkg/api/http/responses"
	"ExpensesTracker/pkg/database"
	"ExpensesTracker/pkg/database/entities"
	"errors"
)

// Source: https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
func UpdateSaving(summaryExpenseId uint, name string, updateSaving requests.CreateSavingRequest) responses.SavingResponse {
	saving := entities.Saving{
		Name:              name,
		ExpensesSummaryId: summaryExpenseId,
		Value:             updateSaving.Value,
	}
	db := database.OpenConnection()
	if db.Model(&saving).Where("name = ?", saving.Name).Updates(&saving).RowsAffected == 0 {
		db.Create(&saving)
	}
	return responses.SavingResponse{
		Name:  saving.Name,
		Value: saving.Value,
	}
}

func DeleteSaving(expensesSummaryId uint, savingId string) error {
	db := database.OpenConnection()
	res := db.Delete(entities.Saving{
		ExpensesSummaryId: expensesSummaryId,
		Name:              savingId,
	})
	if res.RowsAffected == 0 {
		return errors.New("No data found")
	}
	return nil
}
