package initializer

import (
	"github.com/ginniss2022/union/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(
	PG_HOST string,
	PG_PORT string,
	PG_USER string,
	PG_PASS string,
	PG_DB string) {

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", PG_HOST,
	// 	PG_PORT,
	// 	PG_USER,
	// 	PG_PASS,
	// 	PG_DB)

	DB, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=water port=5432"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.User{})
}
