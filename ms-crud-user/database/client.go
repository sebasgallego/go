package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/audited"
	"log"
	"ms-crud-user/entities"
)

type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	//dbURL := "postgres://postgres:admin@localhost:5432/gorm"
	//db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	db, err := gorm.Open("sqlite3", "gorm")
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
	Instance = db
}

func Migrate() {
	Instance.AutoMigrate(&entities.User{})
	log.Println("Database Migration Completed...")
	audited.RegisterCallbacks(Instance)
}
