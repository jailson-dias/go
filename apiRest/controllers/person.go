package controllers

import (
	"encoding/json"
	"go/apiRest/models"

	"net/http"

	"github.com/gorilla/mux"
)

var people []models.Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Person{
		Firstname: "Nome",
		Lastname:  "Last",
	})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
