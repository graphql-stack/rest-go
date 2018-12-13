package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB(fns ...func(db *gorm.DB)) {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}

	for _, fn := range fns {
		fn(db)
	}

	DB = db
}
