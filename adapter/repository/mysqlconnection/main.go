package mysqlconnection

import (
	"fmt"

	"github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

const (
	dbName     = "pismo"
	dbHost     = "localhost"
	dbUser     = "madsonjr"
	dbPassword = "pismo"
	dbPort     = "3306"
)

func connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
