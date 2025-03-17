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
		"host=%s, port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	fmt.Println("Successfully connected to DB")

	// Add automigrate
	err = db.AutoMigrate(&models.User{}, &models.Patient{})
	if err != nil {
		fmt.Println("Failed to migrate DB")
	} else {
		fmt.Println("Successfully Auto Migrated the Database")
		DB = DbInstance{Db: db}
	}

}
