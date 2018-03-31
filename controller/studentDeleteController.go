package controller

import (
	"fmt"
	"net/http"
)

import (
	"../service"
)

func StudentDeleteController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentDeleteController")
	id := r.URL.Query().Get("id")
	fmt.Println(id)

	res := service.DeleteStudent(id)
	if res == false {
		fmt.Println("false")
	} else {
		http.Redirect(w, r, "/student/list", http.StatusSeeOther)
	}

}
