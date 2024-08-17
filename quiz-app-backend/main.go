package main

import (
    "encoding/json"
    "net/http"
    "log"
)

// Question represents a quiz question
type Question struct {
    ID       int      `json:"id"`
    Question string   `json:"question"`
    Options  []string `json:"options"`
    Answer   int      `json:"answer"` // index of the correct option
}

// Sample data
var questions = []Question{
    {1, "What is the capital of France?", []string{"Paris", "London", "Berlin", "Madrid"}, 0},
    {2, "What is 2 + 2?", []string{"3", "4", "5", "6"}, 1},
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(questions)
}

func main() {
    http.HandleFunc("/questions", getQuestions)
    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
