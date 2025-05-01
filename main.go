package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type numbersInt struct {
	A int `json:"A"`
	B int `json:"B"`
}

type numbersFloat struct {
	A float32 `json:"A"`
	B float32 `json:"B"`
}

type resultFloat struct {
	Result float32 `json:"result"`
}

type resultInt struct {
	Result int `json:"result"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	message := `Welcome to the Calculator API!

Usage:
Send POST requests to the following endpoints with JSON body:
{
  "A": 5,
  "B": 3
}

Endpoints:
  /add       -> Returns 8
  /subtract  -> Returns 2
  /multiply  -> Returns 15
  /divide    -> Returns 1.666...

Only basic operations are supported: +, -, *, /
`
	fmt.Fprint(w, message)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}
	var digits numbersInt
	if err := json.NewDecoder(r.Body).Decode(&digits); err != nil {
		log.Println("Unable to parse request body:", err)
		http.Error(w, "Invalid input. Please provide integers A and B.", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(resultInt{Result: digits.A + digits.B})
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}
	var digits numbersInt
	if err := json.NewDecoder(r.Body).Decode(&digits); err != nil {
		log.Println("Unable to parse request body:", err)
		http.Error(w, "Invalid input. Please provide integers A and B.", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(resultInt{Result: digits.A - digits.B})
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}
	var digits numbersInt
	if err := json.NewDecoder(r.Body).Decode(&digits); err != nil {
		log.Println("Unable to parse request body:", err)
		http.Error(w, "Invalid input. Please provide integers A and B.", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(resultInt{Result: digits.A * digits.B})
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}
	var digits numbersFloat
	if err := json.NewDecoder(r.Body).Decode(&digits); err != nil {
		log.Println("Unable to parse request body:", err)
		http.Error(w, "Invalid input. Please provide numbers A and B.", http.StatusBadRequest)
		return
	}
	if digits.B == 0 {
		http.Error(w, "Cannot divide by zero.", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(resultFloat{Result: digits.A / digits.B})
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/subtract", subtractHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	http.HandleFunc("/divide", divideHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
