package configuration

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) *gorm.DB {
	username := config.Get("PG_USERNAME")
	password := config.Get("PG_PASSWORD")
	host := config.Get("PG_HOST")
	port := config.Get("PG_PORT")
	dbName := config.Get("PG_DB_NAME")

	maxPoolOpen, err := strconv.Atoi(config.Get("PG_POOL_MAX_CONN"))
	exception.PanicLogging(err)

	maxPoolIdle, err := strconv.Atoi(config.Get("PG_POOL_IDLE_CONN"))
	exception.PanicLogging(err)

	maxPollLifeTime, err := strconv.Atoi(config.Get("PG_POOL_LIFE_TIME"))
	exception.PanicLogging(err)

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := "host=" + host +
		" port=" + port +
		" user=" + username +
		" password=" + password +
		" dbname=" + dbName +
		" sslmode=disable" // Use `sslmode=disable` for local development

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: loggerDb,
	})
	exception.PanicLogging(err)

	// Configure connection pool
	pgDB, err := db.DB()
	exception.PanicLogging(err)

	pgDB.SetMaxOpenConns(maxPoolOpen)
	pgDB.SetMaxIdleConns(maxPoolIdle)
	pgDB.SetConnMaxLifetime(time.Duration(maxPollLifeTime) * time.Millisecond)

	// Migrate
	// runMigrations(dsn)

	return db
}

// func runMigrations(dsn string) {
// 	conn, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}
// 	defer conn.Close()

// 	migrations := []struct {
// 		up   string
// 		down string
// 	}{
// 		{
// 			up:   "db/migrations/20240831075959_create_family_trees_table.up.sql",
// 			down: "db/migrations/20240831075959_create_family_trees_table.down.sql",
// 		},
// 		{
// 			up:   "db/migrations/20240831080018_create_messages_table.up.sql",
// 			down: "db/migrations/20240831080018_create_messages_table.down.sql",
// 		},
// 		{
// 			up:   "db/migrations/20240831075959_create_family_trees_table.up.sql",
// 			down: "db/migrations/20240831080012_create_users_table.down.sql",
// 		},
// 	}

// 	for _, m := range migrations {
// 		if err := applyMigration(conn, m.up); err != nil {
// 			log.Fatalf("Failed to apply migration: %v", err)
// 		}
// 		// Optionally, apply the down migrations if needed
// 	}
// }

// func applyMigration(db *sql.DB, file string) error {
// 	data, err := os.ReadFile(file)
// 	if err != nil {
// 		return fmt.Errorf("failed to read migration file %s: %w", file, err)
// 	}

// 	_, err = db.Exec(string(data))
// 	if err != nil {
// 		return fmt.Errorf("failed to execute migration file %s: %w", file, err)
// 	}

// 	log.Printf("Successfully applied migration: %s", file)
// 	return nil
// }
