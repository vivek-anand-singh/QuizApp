package database

import "quiz-app-backend/models"

var questions = []models.Question{
    {
        ID:       1,
        Question: "What is the capital of France?",
        Options: []models.Option{
            {ID: 1, Text: "London"},
            {ID: 2, Text: "Berlin"},
            {ID: 3, Text: "Paris"},
            {ID: 4, Text: "Madrid"},
        },
        CorrectAnswer: 3,
    },
	{
		ID:       2,
		Question: "What is the capital of Germany?",
		Options: []models.Option{
			{ID: 1, Text: "London"},
			{ID: 2, Text: "Berlin"},
			{ID: 3, Text: "Paris"},
			{ID: 4, Text: "Madrid"},
		},
		CorrectAnswer: 2,
	},
	{
		ID:			3,
		Question:	"What is the capital of Spain?",
		Options:	[]models.Option{
			{ID: 1, Text: "London"},
			{ID: 2, Text: "Berlin"},
			{ID: 3, Text: "Paris"},
			{ID: 4, Text: "Madrid"},
		},
		CorrectAnswer: 4,
	},

}

func GetQuestions() []models.Question {
    return questions
}

func CheckAnswers(userAnswers []models.UserAnswer) models.QuizResult {
    var result models.QuizResult

    for _, userAnswer := range userAnswers {
        for _, question := range questions {
            if question.ID == userAnswer.QuestionID {
                if question.CorrectAnswer == userAnswer.AnswerID {
                    result.CorrectAnswers++
                } else {
                    result.IncorrectAnswers++
                }
                break
            }
        }
    }

    return result
}