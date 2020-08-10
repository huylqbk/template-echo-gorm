package migrations

import (
	"log"
	"template-echo-gorm/app/models"

	"github.com/jinzhu/gorm"
)

func Migrate(DB *gorm.DB) {
	err := DB.AutoMigrate(&models.User{}, &models.Transaction{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = DB.Model(&models.Transaction{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	log.Println("Migrations Finished")
}
