package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient(dbHost string, dbUser string, dbPassword string, dbName string, dbPort string, dbTimezone string) (*gorm.DB, error) {
	return createClient(dbHost, dbUser, dbPassword, dbName, dbPort, dbTimezone)
}
func checkAndCreateDatabase(DB *gorm.DB, dbName string) {
	var count int64
	result := DB.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count)
	if result.Error != nil {
		panic(result.Error)
	}

	if count == 0 {
		createDatabase(DB, dbName)
	} else {
		fmt.Printf("Database %s already exists\n", dbName)
	}
}

func createDatabase(DB *gorm.DB, dbName string) {
	result := DB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	// DB.Exec(fmt.Sprintf("INSERT INTO migrations (id, name, created_at) VALUES (1, 'create_database_%s', NOW())", dbName))
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("Database %s created\n", dbName)
}

func createClient(dbHost string, dbUser string, dbPassword string, dbName string, dbPort string, dbTimezone string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPassword, "postgres", dbPort, dbTimezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database postgres")
	}

	checkAndCreateDatabase(db, dbName)

	// dsnNew := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)
	dsnNew := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbTimezone)

	db, err = gorm.Open(postgres.Open(dsnNew), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database with new DSN")
	}

	return db, nil
}
