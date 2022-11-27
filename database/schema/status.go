package schema

import "database/sql/driver"

type TaskStatus string

const (
	Pending  TaskStatus = "pending"
	Complete TaskStatus = "complete"
)

func (r *TaskStatus) Scan(value interface{}) error {
	*r = TaskStatus(value.([]byte))
	return nil
}

func (r TaskStatus) Value() (driver.Value, error) {
	return string(r), nil
}

type Status struct {
	ID   uint       `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name TaskStatus `json:"name" gorm:"unique;not null"`
}
