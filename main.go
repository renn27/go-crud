package main

import (
	"net/http"

	"github.com/renn27/go-crud/controllers/usercontroller"
)

func main() {
	http.HandleFunc("/", usercontroller.Index)
	http.HandleFunc("/user", usercontroller.Index)
	http.HandleFunc("/user/index", usercontroller.Index)
	http.HandleFunc("/user/add", usercontroller.Add)
	http.HandleFunc("/user/edit", usercontroller.Edit)
	http.HandleFunc("/user/delete", usercontroller.Delete)

	http.ListenAndServe(":5000", nil)
}
