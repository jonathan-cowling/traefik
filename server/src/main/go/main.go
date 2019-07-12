package main

import (
	"net/http"
	"os"
	"strconv"
	"math/rand"
	"math"
	"html/template"
)

func main() {

	TEMPLATES := template.Must(template.ParseGlob("templates/*.go.html"))
	T_200 := "200.go.html"
	T_400 := "400.go.html"

	envSuccessThreshold := os.Getenv("SUCCESS_THRESHOLD")

	if (envSuccessThreshold == "") {
		panic("SUCCESS_THRESHOLD not in environment")
	}

	SUCCESS_THRESHOLD, _ := strconv.ParseFloat(envSuccessThreshold, 64)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		randomNumber := rand.NormFloat64()
		successChance := randomNumber - math.Floor(randomNumber)

		if (successChance > SUCCESS_THRESHOLD) {
			w.WriteHeader(http.StatusBadRequest)
			TEMPLATES.ExecuteTemplate(w, T_400, nil)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		TEMPLATES.ExecuteTemplate(w, T_200, nil)
	})

	http.ListenAndServe(":8080", nil)
}