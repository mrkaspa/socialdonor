package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

// it can be used for jobs
var Gdb gorm.DB

// init db
func InitDB() {
	// open db
	fmt.Println("*** INIT DB ***")
	// connString := revel.Config.StringDefault("db.conn", "")
	var connString string
	if revel.Config.BoolDefault("mode.test", false) || revel.Config.BoolDefault("mode.dev", false) {
		connString = os.Getenv("MYSQL_DB")
	} else {
		connString = revel.Config.StringDefault("db.prod", "user:password@tcp(mariadb:3306)/db_name")
	}
	fmt.Printf("Connecting to >> %s", connString)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		fmt.Println("Unable to connect to the database")
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	db.DB().Ping()
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Request{})
	// uniquie index if need
	db.Model(&User{}).AddUniqueIndex("idx_email", "email")
	db.Model(&User{}).AddUniqueIndex("idx_fb_id", "fb_id")
	Gdb = db
}
