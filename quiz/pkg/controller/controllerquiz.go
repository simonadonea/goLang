package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"quiz/pkg/model"
	"quiz/pkg/resources"
)

func QuizPage(w http.ResponseWriter, r *http.Request) {
	temp := template.New("")
	temp.Parse(resources.QuizForm)
	fmt.Println(r.Method)

	quiz := model.GetAllQuestions()
	res, _ := json.Marshal(quiz)
	w.Write(res)
	// fmt.Println(quiz)
	temp.Execute(w, quiz)
}
