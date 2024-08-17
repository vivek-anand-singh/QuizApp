package main

import (
    "log"
    "net/http"
    "quiz-app-backend/routes"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/questions", routes.GetQuestions).Methods("GET")
    r.HandleFunc("/submit", routes.SubmitAnswers).Methods("POST")

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
        AllowCredentials: true,
    })

    handler := c.Handler(r)

    log.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}