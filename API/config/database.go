package initializers
import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )
  
  var DB *gorm.DB

func ConnectToDatabase(
	 PG_HOST string, 
	 PG_PORT string, 
	 PG_USER string, 
	 PG_PASS string, 
	 PG_DB string,  ) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", PG_HOST, 
	 PG_PORT, 
	 PG_USER, 
	 PG_PASS, 
	 PG_DB)

	 DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	 if err != nil {
		return nil, err
	}
	return DB, err
}