package studentController

import (
	"fmt"
	"net/http"
)

import (
	"../../service/studentService"
)

func StudentDeleteController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StudentDeleteController")
	id := r.URL.Query().Get("id")
	fmt.Println(id)

	res := studentService.DeleteStudent(id)
	if res == false {
		fmt.Println("false")
	} else {
		http.Redirect(w, r, "/student/list", http.StatusSeeOther)
	}

}
