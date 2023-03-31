package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/person", getPerson)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	person := Person{Name: "John Doe", Age: 30}
	json.NewEncoder(w).Encode(person)
}
