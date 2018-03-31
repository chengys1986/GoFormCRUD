package main

import (
	"log"
	"net/http"
)

import (
	"./controller"
)

func main() {

	// Routing
	http.HandleFunc("/", controller.HomeController)
	http.HandleFunc("/student/new", controller.StudentNewController)
	http.HandleFunc("/student/list", controller.StudentListController)
	http.HandleFunc("/student/form", controller.StudentFormController)
	http.HandleFunc("/student/edit", controller.StudentEditController)
	http.HandleFunc("/student/delete", controller.StudentDeleteController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
