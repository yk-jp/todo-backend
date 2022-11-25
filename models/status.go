package models

type Status struct {
	ID   uint   `json:"id" gorm:"primaryKey,AUTO_INCREMENT"`
	Name string `json:"name" gorm:"not null"`
}
