package database

import (
	"ExpensesTracker/pkg/api/http/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunMigrations() {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=ExpensesTracker port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&model.ExpensesSummary{},
		&model.Expense{},
		&model.Income{},
		&model.Saving{})
	if err != nil {
		panic("Failed to run migrations")
	}
}
