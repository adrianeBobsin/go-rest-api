package database

import (
	"jwt-authentication-golang/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDBPostegreSQL() {
	dns := "host=localhost user=admin password=admin dbname=admin port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dns))

	if err != nil {
		log.Panic("Connection Database fail.")
	}
	log.Println("Connected to Database!")
}

func ConnectionDBMySQL() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Connection Database fail.")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	DB.AutoMigrate(&models.Client{})
	log.Println("Database Migration Completed!")
}

func CloseConnection() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
