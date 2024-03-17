package database

import (
	"Simple-REST-ful-Fiber/config"
	"Simple-REST-ful-Fiber/internal/repository"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	var err error
	p := config.Config("DBPORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DBUSER"),
		config.Config("DBPASS"),
		config.Config("DBHOST"),
		port,
		config.Config("DBNAME"),
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	fmt.Println("Connect Opened to Database")

	DB.AutoMigrate(&repository.Product{})
	return DB
}