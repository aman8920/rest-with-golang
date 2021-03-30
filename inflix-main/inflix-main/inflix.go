package main

import (
	"fmt"

	"inflix-main/config"

	"inflix-main/model"

	"inflix-main/rest"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// Creating a connection to the database
	config.DB, err = gorm.Open("mysql", config.DbConnectionString(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Error Connecting to Inflix DB: ", err)
	}

	defer config.DB.Close()

	// run the migrations: album struct
	config.DB.AutoMigrate(&model.Album{})

	//setup routes
	r := rest.SetupRouter()

	// running
	r.Run()

}
