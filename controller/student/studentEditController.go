package studentController

import (
	"fmt"
	"net/http"
	"strconv"
)

import (
	"../../model"
	"../../service/studentService"
	"../../service/templateService"
)

func StudentEditController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentEditController")

	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		err, student := studentService.GetStudent(id)
		if err != nil {
			fmt.Println("oppos")
		} else {
			outStr := templateService.GetStudentEditPage(student)

			w.Header().Set("Content-Type", "text/html; charset-utf-8")
			w.Write([]byte(outStr))
		}
	} else if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("studentid"))
		if err == nil {
			res := studentService.UpdateStudent(model.Student{StudentID: id, Name: r.FormValue("name"), Description: r.FormValue("description")})
			if res == false {
				fmt.Println("false")
			} else {
				http.Redirect(w, r, "/student/list", http.StatusSeeOther)
			}
		}
	}
}
