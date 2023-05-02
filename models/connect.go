// models/connect.go

package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDB() {
	
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "dbdUser1:hY3aW0wR4zoP4sV9pG7w@tcp(kairet.ru:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
                panic("Failed to connect to database!")
        }

	err = database.AutoMigrate(&User{})
        if err != nil {
                return
        }

        DB = database
}
