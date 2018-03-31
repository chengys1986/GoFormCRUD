package controller

import (
	"fmt"
	"net/http"
)

import (
	"../service"
)

func StudentFormController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentFormController")

	id := r.URL.Query().Get("id")
	err, student := service.GetStudent(id)
	if err != nil {
		fmt.Println("oppos")
	} else {
		outStr := service.GetStudentFormPage(student)

		w.Header().Set("Content-Type", "text/html; charset-utf-8")
		w.Write([]byte(outStr))
	}
}
