package models

type Task struct {
	ID           uint   `json:"id" gorm:"primaryKey,AUTO_INCREMENT"`
	Title        string `json:"title" gorm:"not null"`
	Status_Refer int    `json:"status_id" gorm:"not null"`
	Status       Status `gorm:"foreignKey:Status_Refer"`
}
