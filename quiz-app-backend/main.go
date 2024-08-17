package main

import (
    "log"
    "net/http"
    "quiz-app-backend/routes"
    "github.com/rs/cors"
)

func main() {
    // Initialize a new router (you might be using a different router, like gorilla/mux, in your actual routes)
    mux := http.NewServeMux()
    mux.HandleFunc("/questions", routes.GetQuestions)

    // Set up CORS options
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Adjust this to your needs
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
        AllowCredentials: true,
    })

    // Wrap your router with the CORS middleware
    handler := c.Handler(mux)

    log.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
