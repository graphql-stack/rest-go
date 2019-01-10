package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zcong1993/rest-go/model"
	"log"
	"os"
)

var DB *gorm.DB

func init() {
	InitDB(func(db *gorm.DB) {
		db.AutoMigrate(new(model.User))
		db.AutoMigrate(new(model.Token))
		db.AutoMigrate(new(model.Book))
	})
}

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
