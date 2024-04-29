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

	InitData()

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

func InitData() error {
	sampleData := []models.Customer{
		{Name: "John Doe", Age: 30},
		{Name: "Jane Smith", Age: 25},
		{Name: "Alice Johnson", Age: 40},
		{Name: "Bob Brown", Age: 35},
		{Name: "Emma Davis", Age: 28},
		{Name: "Michael Wilson", Age: 45},
		{Name: "Olivia Taylor", Age: 33},
		{Name: "William Martinez", Age: 39},
		{Name: "Sophia Garcia", Age: 31},
		{Name: "James Rodriguez", Age: 37},
	}

	for _, data := range sampleData {
		if err := DB.Create(&data).Error; err != nil {
			return err
		}
	}

	return nil
}
