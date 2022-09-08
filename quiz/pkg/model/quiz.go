package model

import (
	"fmt"
	"quiz/pkg/config"
)

type Quiz []QuizQuestion

type QuizQuestion struct {
	ID              uint     `gorm:"primaryKey"json:"id"`
	Question        string   `json:"question"`
	QuestionAnswers []Answer `json:"questionAnswers"`
	UserAnswer      string   `json:"useranswer"`
}

type Answer struct {
	ID             uint   `gorm:"primaryKey"json:"id"`
	QuizQuestionID uint   `json:"questionId"`
	Text           string `json:"text"`
	IsCorrect      bool   `json:"iscorrect"`
}

func init() {
	db = config.Connect()
	db.AutoMigrate(&Answer{})
	db.AutoMigrate(&QuizQuestion{})
}

// func (q *QuizQuestion) CreateQuestion() *QuizQuestion{
// 	db.NewRecord(q)
// 	db.Create(&q)
// 	return q
// }

func GetAllQuestions() []QuizQuestion {
	var quiz Quiz
	var answers []Answer
	db.Find(&quiz)
	db.Find(&answers)

	fmt.Println("Questions:")
	for i := range quiz {
		quiz[i].QuestionAnswers = make([]Answer, 0)
		fmt.Printf("%d: %v\n", i, quiz[i])
	}

	fmt.Println("Answers:")
	for i, v := range answers {
		fmt.Printf("%d: %v\n", i, v)
	}

	for i := range answers {
		a := answers[i]
		question, err := quiz.findQuestion(a.QuizQuestionID)
		// fmt.Println(question)
		if err != nil {
			panic(err)
		}
		question.QuestionAnswers = append(question.QuestionAnswers, a)
	}

	fmt.Println("Questions:")
	for i := range quiz {
		fmt.Printf("%d: %v\n", i, quiz[i])
	}

	return quiz
}

func (quiz Quiz) findQuestion(id uint) (*QuizQuestion, error) {
	for i := range quiz {
		if quiz[i].ID == id {
			return &quiz[i], nil
		}
	}

	return nil, fmt.Errorf("No question with id %d", id)
}
