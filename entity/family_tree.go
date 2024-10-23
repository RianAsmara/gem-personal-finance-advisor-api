package entity

import (
	"time"

	"github.com/google/uuid"
)

type FamilyTree struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// Users       []User `gorm:"foreignKey:FamilyTreeID"`
}

func (FamilyTree) TableName() string {
	return "family_trees"
}
