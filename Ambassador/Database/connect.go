package database

import (
	models "ambassador/Models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_USER_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database_name := os.Getenv("MYSQL_DATABASE_NAME")

	// user:pass@tcp(127.0.0.1:3306)/dbname
	database, err := gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, password, host, port, database_name)), &gorm.Config{})

	if err != nil {
		panic("couldn't connect to database")
	}

	fmt.Println("Connected to database")

	DB = database

	database.AutoMigrate(&models.User{}, &models.Product{})

}
