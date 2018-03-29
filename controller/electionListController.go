package controller

import (
	"fmt"
	"net/http"
)

func electionListController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi %s", r.URL.Path[1:])
}
