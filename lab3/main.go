package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Задание 1: Обработка query-параметров
func handleGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	ageStr := r.URL.Query().Get("age")

	if name == "" || ageStr == "" {
		http.Error(w, "Необходимо передать параметры name и age", http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, "Неверный формат параметра age", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("Меня зовут %s, мне %d лет", name, age)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}

// Задание 2: Маршрутизация и арифметика
func handleArithmetic(w http.ResponseWriter, r *http.Request) {

	op := r.URL.Path[1:]
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	if aStr == "" || bStr == "" {
		http.Error(w, "Необходимо передать параметры a и b", http.StatusBadRequest)
		return
	}

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		http.Error(w, "Неверный формат параметра a", http.StatusBadRequest)
		return
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		http.Error(w, "Неверный формат параметра b", http.StatusBadRequest)
		return
	}

	var result float64
	var errOp error
	switch op {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			http.Error(w, "Деление на ноль", http.StatusBadRequest)
			return
		}
		result = a / b
	default:
		errOp = fmt.Errorf("неизвестная операция")
		http.Error(w, errOp.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("%f", result)))
}

// Задание 3: Обработка JSON
func handleCharCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST запросы", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Text string `json:"text"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	counts := make(map[string]int)
	for _, char := range requestBody.Text {
		counts[string(char)]++
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(counts)
}

func main() {
	http.HandleFunc("/greet", handleGreeting)
	http.HandleFunc("/add", handleArithmetic)
	http.HandleFunc("/sub", handleArithmetic)
	http.HandleFunc("/mul", handleArithmetic)
	http.HandleFunc("/div", handleArithmetic)
	http.HandleFunc("/count", handleCharCount)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
