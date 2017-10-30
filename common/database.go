package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../gorm.db")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open("sqlite3", "./../gorm_test.db")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	DB = test_db
	return DB
}

func TestDBFree(test_db *gorm.DB) error {
	test_db.Close()
	err := os.Remove("./../gorm_test.db")
	return err
}

func GetDB() *gorm.DB {
	return DB
}
