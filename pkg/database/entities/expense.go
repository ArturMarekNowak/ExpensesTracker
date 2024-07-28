package entities

type Expense struct {
	ExpensesSummaryId uint    `gorm:"primaryKey"`
	Name              string  `gorm:"primaryKey"`
	Value             float32 `gorm:"type:numeric(10,2)"`
}
