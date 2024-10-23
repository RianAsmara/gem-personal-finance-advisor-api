package entity

import "time"

type Role struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []User `gorm:"many2many:user_roles"`
}

func (Role) TableName() string {
	return "roles"
}
