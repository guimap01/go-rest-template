package database

import (
	"go-api-rest/src/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Setup() {
	host := "host"
	port := "port"
	dbname := "dbname"
	user := "user"
	password := "password"

	if db, err := gorm.Open(
		"postgres",
		"host="+host+
			" port="+port+
			" user="+user+
			" dbname="+dbname+
			" sslmode=disable password="+password); err != nil {
		log.Fatal(err)
	} else {
		db.LogMode(false)
		db.AutoMigrate([]models.Book{})
		DB = db
	}
}

func GetDB() *gorm.DB {
	return DB
}

func ClearTable(tableName string) {
	DB.Exec("DELETE FROM %s", tableName)
	DB.Exec("ALTER SEQUENCE %s_id_seq RESTART WITH 1", tableName)
}
