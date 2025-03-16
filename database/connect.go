package database

import (
	"CarepluseBackend/config"
	"CarepluseBackend/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")

	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		panic("Failed to parse DB Port")
	}

	dsn := fmt.Sprintf(
		"host=db, port=%d user=%s password=%s dbname=%s sslmode=disable", port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"),
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	fmt.Println("Successfully connected to DB")

	// Add automigrate
	err = DB.AutoMigrate(&models.User{}, &models.Patient{})
	if err != nil {
		fmt.Println("Failed to migrate DB")
	}
	fmt.Println("Successfully Auto Migrated the Database")
}
