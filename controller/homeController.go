package controller

import (
	"fmt"
	"net/http"
)

import (
	"../service"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home controller")
	w.Write([]byte(service.GetHomePage()))
}
