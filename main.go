package main

import (
	"fmt"
	"log"
	"net/http"
)

import (
	"./controller"
	"./service"
	//	"githut.com/chengys1986/GoFormCRUD/service/dbservice"
)

func main() {
	fmt.Println("Starting up")

	// init db
	db, err := service.InitDB()
	if err != nil {
		fmt.Printf("DB Error: $s\n", err.Error())
		return
	}
	defer service.CloseDB(db)

	// Routing
	http.HandleFunc("/", controller.HomeController)
	http.HandleFunc("/student/new", controller.StudentNewController)
	http.HandleFunc("/student/list", controller.StudentListController)
	http.HandleFunc("/student/form", controller.StudentFormController)
	http.HandleFunc("/student/edit", controller.StudentEditController)
	http.HandleFunc("/student/delete", controller.StudentDeleteController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
