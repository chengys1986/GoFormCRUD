package studentController

import (
	"fmt"
	"net/http"
)

import (
	"../../service/studentService"
	"../../service/templateService"
)

func StudentListController(w http.ResponseWriter, r *http.Request) {
	students := studentService.GetStudents()
	if students == nil {
		fmt.Println("oops")
	} else {
		outStr := templateService.GetStudentListPage(students)
		if outStr != "" {
			w.Write([]byte(outStr))
		} else {
			fmt.Println("opps")
		}
	}
}
