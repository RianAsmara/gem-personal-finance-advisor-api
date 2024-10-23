package migrations

import (
	"github.com/RianAsmara/personal-finance-advisor-api/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.FamilyTree{},
		&entity.Message{},
		&entity.Role{},
	)
}
