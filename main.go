package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	Method string
	Value  string
}

var i int = 10

func add(w http.ResponseWriter, r *http.Request) {
	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	i = i + value
	var result string = strconv.Itoa(i)
	response := Response{"sub", result}
	js, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func sub(w http.ResponseWriter, r *http.Request) {
	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	i = i - value
	var result string = strconv.Itoa(i)
	response := Response{"sub", result}
	js, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleRequests() {
	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	handleRequests()
}
