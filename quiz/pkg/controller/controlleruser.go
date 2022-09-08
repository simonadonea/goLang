package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"quiz/pkg/model"
	"quiz/pkg/resources"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	temp := executeTemplate(resources.Index, w, r)
	temp.Execute(w, nil)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	tmpl := executeTemplate(resources.Signup, w, r)

	user := model.User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	user.CreateUser()
	fmt.Println(user)

	tmpl.Execute(w, struct{ Success bool }{true})
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	temp := executeTemplate(resources.Signup, w, r)
	temp.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := executeTemplate(resources.Login, w, r)
	username := r.FormValue("username")
	password := r.FormValue("password")

	dbUser := model.GetUserByUsername(username)
	fmt.Println(username, password)
	if dbUser.Username == username && dbUser.Password == password {
		tmpl.Execute(w, struct {
			Success bool
			Fail    bool
		}{true, false})
	} else {
		fmt.Println("failed login")

		tmpl.Execute(w, struct {
			Success bool
			Fail    bool
		}{false, true})
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	temp := executeTemplate(resources.Login, w, r)
	temp.Execute(w, nil)
}

func executeTemplate(resource string, w http.ResponseWriter, r *http.Request) *template.Template {
	temp := template.New("")
	temp.Parse(resource)
	return temp
}
