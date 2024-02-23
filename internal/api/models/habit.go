package models

type Habit struct {
	BaseModel
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	UserID      uint
}

// TableName sets the table name for the Habit model.
func (Habit) TableName() string {
	return "habits"
}
