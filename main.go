package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

type StatusValue struct {
	WaterValue  int
	WaterStatus string
	WindValue   int
	WindStatus  string
}

func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	statusValue := GetStatusValue()

	var t, err = template.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(w, statusValue)
}

func randomIntGen(start, end int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start) + start
}

func GetStatusValue() StatusValue {
	var statusValue StatusValue

	statusValue.WaterValue = randomIntGen(1, 100)
	time.Sleep(1)
	statusValue.WindValue = randomIntGen(1, 100)

	if statusValue.WaterValue <= 5 {
		statusValue.WaterStatus = "Aman"
	} else if statusValue.WaterValue >= 6 && statusValue.WaterValue <= 8 {
		statusValue.WaterStatus = "Siaga"
	} else if statusValue.WaterValue > 8 {
		statusValue.WaterStatus = "Bahaya"
	} else {
		statusValue.WaterStatus = "Error"
	}

	if statusValue.WindValue <= 6 {
		statusValue.WindStatus = "Aman"
	} else if statusValue.WindValue >= 7 && statusValue.WindValue <= 15 {
		statusValue.WindStatus = "Siaga"
	} else if statusValue.WindValue > 15 {
		statusValue.WindStatus = "Bahaya"
	} else {
		statusValue.WindStatus = "Error"
	}

	return statusValue
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/get-status", getStatusHandler)

	fmt.Println("running at localhost:8080")
	http.ListenAndServe(":8080", router)
}
