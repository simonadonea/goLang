package routes

import (
	"quiz/pkg/controller"

	"github.com/gorilla/mux"
)

var RegisterUsersRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controller.MainPage).Methods("GET")
	router.HandleFunc("/signup", controller.SignupPage).Methods("GET")
	router.HandleFunc("/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/login", controller.LoginPage).Methods("GET")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/quiz", controller.QuizPage).Methods("GET")
}
