package main

import (
	"net/http"

	"github.com/renn27/go-crud/controllers/usercontroller"
)

func main() {
	http.HandleFunc("/", usercontroller.Index)
	http.HandleFunc("/", usercontroller.Index)
}
