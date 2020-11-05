package models

// Todo is model for todo row
type Todo struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// TableName returns table name
func (b *Todo) TableName() string {
	return "todo"
}
