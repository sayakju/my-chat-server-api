package handler

import "github.com/gorilla/mux"

func NewRouter(
	peopleHandler PeopleHandler,
) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/people", peopleHandler.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", peopleHandler.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", peopleHandler.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", peopleHandler.DeletePerson).Methods("DELETE")
	return router
}