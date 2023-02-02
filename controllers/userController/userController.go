package usercontroller

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/renn27/go-crud/entities"
)

func Index(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/user/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/user/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var user entities.User
		user.idUser = request.Form.Get("id_user")
		user.nama = request.Form.Get("nama")
		user.alamat = request.Form.Get("alamat")
		user.no_telp = request.Form.Get("no_hp")

		fmt.Println(user)
	}
}
func Edit(response http.ResponseWriter, request *http.Request) {

}
func Delete(response http.ResponseWriter, request *http.Request) {

}
