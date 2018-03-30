package controller

import (
	"fmt"
	"net/http"
)

import (
	"../service"
)

func StudentListController(w http.ResponseWriter, r *http.Request) {
	students := service.GetStudents()
	if students == nil {
		fmt.Println("oops")
	} else {
		outStr := service.GetStudentListPage(students)
		if outStr != "" {
			w.Write([]byte(outStr))
		} else {
			fmt.Println("opps")
		}
	}
}
