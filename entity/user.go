package entity

import (
	"time"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	IsActive bool   `gorm:"default:false"`
	// FamilyTreeID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	// FamilyTree FamilyTree `gorm:"foreignKey:FamilyTreeID"`
	Roles []Role `gorm:"many2many:users_roles;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (User) TableName() string {
	return "users"
}
