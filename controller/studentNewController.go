package controller

import (
	"fmt"
	"net/http"
)

import (
	"../model"
	"../service"
)

func StudentNewController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		outStr := service.GetStudentNewPage()
		if outStr == "" {
			fmt.Println("opps")
		} else {
			w.Header().Set("Content-Type", "text/html; charset-utf-9")
			w.Write([]byte(outStr))
		}
	} else {

		res := service.NewStudent(model.Student{Name: r.FormValue("name"), Description: r.FormValue("description")})
		if res == false {
			fmt.Println("false")
		} else {
			http.Redirect(w, r, "/student/list", http.StatusSeeOther)
		}
	}
}
