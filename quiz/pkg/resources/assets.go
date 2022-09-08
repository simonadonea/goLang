package resources

import (
	_ "embed"
)

//go:embed index.html
var Index string

//go:embed login.html
var Login string

//go:embed signup.html
var Signup string

//go:embed quiz.html
var QuizForm string
