package main

import (
	"fmt"
	"mvc_pattern/Config"
	Models "mvc_pattern/Models/User"
	"mvc_pattern/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open(
		"mysql",
		Config.DbURL(Config.BuildDBConfig()),
	)

	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()

	// running
	r.Run()
}
