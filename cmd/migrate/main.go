package main

import (
	"log"

	"github.com/RianAsmara/personal-finance-advisor-api/configuration"
	"github.com/RianAsmara/personal-finance-advisor-api/db/migrations"
)

func main() {
	config := configuration.New()
	db := configuration.NewDatabase(config)

	err := migrations.RunMigrations(db)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	log.Println("Migrations completed successfully")
}
