package configs

import (
	"log"
	"github.com/Safwanseban/interview-patient/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connecttodb() {
	var err error
	dsn := "host=localhost user=safwan password=Safwan@123 dbname=gorm port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Println("err connecting to db", err)
	}
	Db.AutoMigrate(&models.Patient{})

}




