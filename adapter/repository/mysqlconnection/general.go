package mysqlconnection

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbName     = "pismo"
	dbHost     = "localhost"
	dbUser     = "madsonjr"
	dbPassword = "pismo"
	dbPort     = "3306"
)

func connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
