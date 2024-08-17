package routes

import (
    "encoding/json"
    "net/http"
    "quiz-app-backend/models"
)

var questions = []models.Question{
    {ID: 1, Question: "What is the capital of France?", Answer: "Paris"},
    {ID: 2, Question: "Who wrote 'To Kill a Mockingbird'?", Answer: "Harper Lee"},
    {ID: 3, Question: "What is the largest planet in our solar system?", Answer: "Jupiter"},
    {ID: 4, Question: "What year did the Titanic sink?", Answer: "1912"},
    {ID: 5, Question: "What is the hardest natural substance on Earth?", Answer: "Diamond"},
    {ID: 6, Question: "Who painted the Mona Lisa?", Answer: "Leonardo da Vinci"},
    {ID: 7, Question: "What is the smallest country in the world?", Answer: "Vatican City"},
    {ID: 8, Question: "What element does 'O' represent on the periodic table?", Answer: "Oxygen"},
    {ID: 9, Question: "What is the chemical symbol for gold?", Answer: "Au"},
    {ID: 10, Question: "Who is known as the father of modern physics?", Answer: "Albert Einstein"},
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(questions)
}
