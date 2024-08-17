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
        ID:       3,
        Question: "What is the capital of Spain?",
        Options: []models.Option{
            {ID: 1, Text: "London"},
            {ID: 2, Text: "Berlin"},
            {ID: 3, Text: "Paris"},
            {ID: 4, Text: "Madrid"},
        },
        CorrectAnswer: 4,
    },
    {
        ID:       4,
        Question: "What is the largest planet in our solar system?",
        Options: []models.Option{
            {ID: 1, Text: "Earth"},
            {ID: 2, Text: "Mars"},
            {ID: 3, Text: "Jupiter"},
            {ID: 4, Text: "Saturn"},
        },
        CorrectAnswer: 3,
    },
    {
        ID:       5,
        Question: "Who wrote 'To Kill a Mockingbird'?",
        Options: []models.Option{
            {ID: 1, Text: "Harper Lee"},
            {ID: 2, Text: "Mark Twain"},
            {ID: 3, Text: "Ernest Hemingway"},
            {ID: 4, Text: "J.K. Rowling"},
        },
        CorrectAnswer: 1,
    },
    {
        ID:       6,
        Question: "What is the chemical symbol for gold?",
        Options: []models.Option{
            {ID: 1, Text: "Au"},
            {ID: 2, Text: "Ag"},
            {ID: 3, Text: "Pb"},
            {ID: 4, Text: "Fe"},
        },
        CorrectAnswer: 1,
    },
    {
        ID:       7,
        Question: "In which year did the Titanic sink?",
        Options: []models.Option{
            {ID: 1, Text: "1905"},
            {ID: 2, Text: "1912"},
            {ID: 3, Text: "1923"},
            {ID: 4, Text: "1931"},
        },
        CorrectAnswer: 2,
    },
    {
        ID:       8,
        Question: "What is the hardest natural substance on Earth?",
        Options: []models.Option{
            {ID: 1, Text: "Gold"},
            {ID: 2, Text: "Iron"},
            {ID: 3, Text: "Diamond"},
            {ID: 4, Text: "Platinum"},
        },
        CorrectAnswer: 3,
    },
    {
        ID:       9,
        Question: "Who painted the Mona Lisa?",
        Options: []models.Option{
            {ID: 1, Text: "Vincent van Gogh"},
            {ID: 2, Text: "Pablo Picasso"},
            {ID: 3, Text: "Leonardo da Vinci"},
            {ID: 4, Text: "Claude Monet"},
        },
        CorrectAnswer: 3,
    },
    {
        ID:       10,
        Question: "What is the capital of Japan?",
        Options: []models.Option{
            {ID: 1, Text: "Seoul"},
            {ID: 2, Text: "Beijing"},
            {ID: 3, Text: "Tokyo"},
            {ID: 4, Text: "Bangkok"},
        },
        CorrectAnswer: 3,
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
