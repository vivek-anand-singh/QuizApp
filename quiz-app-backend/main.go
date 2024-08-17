package main

import (
    "net/http"
    "quiz-app-backend/routes"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func CreateHandler() http.Handler {
    r := mux.NewRouter()

    r.HandleFunc("/questions", routes.GetQuestions).Methods("GET")
    r.HandleFunc("/submit", routes.SubmitAnswers).Methods("POST")

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
        AllowCredentials: true,
    })

    return c.Handler(r)
}

func main() {
    http.ListenAndServe(":8080", CreateHandler())
}
