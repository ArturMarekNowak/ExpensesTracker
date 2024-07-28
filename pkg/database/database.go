package database

import (
	"ExpensesTracker/pkg/database/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
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

	db, err := gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
