package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Person struct
type Person struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Age    int32   `json:"age"`
	Hobbie *Hobbie `json:"hobbie"`
}

// Hobbie struct
type Hobbie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Init persons var as a slice Person struct
var persons []Person

func main() {
	fmt.Println("Hola")

	// Init router
	r := mux.NewRouter()

	persons = append(persons, Person{ID: "1", Name: "Joe", Age: 13, Hobbie: &Hobbie{Title: "Play guitar", Description: "only play"}})

	// Route handles & Endopoints
	r.HandleFunc("/persons", getPersons).Methods("GET")          //🆗
	r.HandleFunc("/persons/{id}", getPerson).Methods("GET")      //🆗
	r.HandleFunc("/persons", createPerson).Methods("POST")       //🆗
	r.HandleFunc("/persons/{id}", updatePerson).Methods("PUT")   //🆗
	r.HandleFunc("/persons/{id}", deletePerson).Methods("DELTE") //🆗

	// Start Server
	log.Fatal(http.ListenAndServe(":5000", r))

}

// Get all persons
func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

// Get single Person
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through persons and find one with the id from the params
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Add new Person
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = strconv.Itoa(rand.Intn(100000000))
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)
}

// Update Person
func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Application-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			var person Person
			_ = json.NewDecoder(r.Body).Decode(&person)
			person.ID = params["id"]
			persons = append(persons, person)
			json.NewEncoder(w).Encode(person)
			return
		}
	}
}

// Delete Person
func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)
}
