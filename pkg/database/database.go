package database

import (
	"ExpensesTracker/pkg/database/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunMigrations() {

	db := OpenConnection()

	err := db.AutoMigrate(&entities.ExpensesSummary{},
		&entities.Expense{},
		&entities.Income{},
		&entities.Saving{})
	if err != nil {
		panic("Failed to run migrations")
	}
}

func OpenConnection() *gorm.DB {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=ExpensesTracker port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
