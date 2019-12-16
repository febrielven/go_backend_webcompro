package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	db *gorm.DB
)

func Connect(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s Database", Dbdriver)
		log.Fatal("this is the error", err)
	} else {
		fmt.Printf("connect to database successs")
	}

}

func GetDB() *gorm.DB {
	return db
}
