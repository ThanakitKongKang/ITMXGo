package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"itmxgo/models"
)

var DB *gorm.DB

const DatabaseType = "sqlite3"
const DatabasePath = "db/itmxgo.db"

func Init() error {
	var err error
	DB, err = gorm.Open(DatabaseType, DatabasePath)
	DB.AutoMigrate(&models.Customer{})

	if err := ResetDB(); err != nil {
		return err
	}

	return err
}

func SetupTestDB() {
	testDatabasePath := "itmxgo_test.db"
	var err error
	DB, err = gorm.Open(DatabaseType, testDatabasePath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	// Perform auto migration
	if err := DB.AutoMigrate(&models.Customer{}).Error; err != nil {
		fmt.Println("Error performing auto migration:", err)
		return
	}
}

func ResetDB() error {
	// Truncate the Customer table
	if err := DB.Exec("DELETE FROM customers").Error; err != nil {
		return err
	}

	// Reset the auto-increment sequence
	if err := DB.Exec("DELETE FROM sqlite_sequence WHERE name=?", "customers").Error; err != nil {
		return err
	}

	return nil
}
