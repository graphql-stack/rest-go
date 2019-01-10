package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/zcong1993/rest-go/model"
	"log"
	"os"
)

var ORM *gorm.DB

func InitDB(fns ...func(db *gorm.DB)) {
	db, err := gorm.Open("postgres", os.Getenv("PG_URL"))
	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(50)

	for _, fn := range fns {
		fn(db)
	}

	if os.Getenv("ENV") == "debug" {
		db.LogMode(true)
	}

	ORM = db
}

func init() {
	InitDB(func(db *gorm.DB) {
		db.AutoMigrate(new(model.User))
		db.AutoMigrate(new(model.Token))
		db.AutoMigrate(new(model.Post))
		db.AutoMigrate(new(model.Comment))
	})
}
