package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go/apiRest/controllers"
)

func main() {
	router := mux.NewRouter()
	// people = append(people, models.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	// people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	router.HandleFunc("/contato", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", controllers.GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", controllers.CreatePerson).Methods("POST")
	// router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
