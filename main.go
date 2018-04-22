package main

import (
	"fmt"
)

import (
	"./service/config"
	"./service/db"
	"./service/webservice"
)

func main() {
	fmt.Println("Starting up")

	// init config
	configErr := config.InitConfig()
	if configErr != nil {
		fmt.Printf("Config Error: %s\n", configErr.Error())
		return
	}

	// init db
	conn, err := db.InitDB()
	if err != nil {
		fmt.Printf("DB Error: %s\n", err.Error())
		return
	}
	defer db.CloseDB(conn)

	// Routing
	webservice.StartWebServer()
}
