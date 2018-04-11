package homeController

import (
	"fmt"
	"net/http"
)

import (
	"../../service/templateService"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home controller")
	w.Write([]byte(templateService.GetHomePage()))
}
