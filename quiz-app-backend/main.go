package main

import (
    "log"
    "net/http"
    "quiz-app-backend/routes"
)

func main() {
    http.HandleFunc("/questions", routes.GetQuestions)

    log.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
