//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
)

type RequestError struct {
	Error 	string		`json:"error"`
	Data 	interface{}	`json:"data,omitempty"`
}

func Answer(data interface{}, w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	apiCtrl := NewAPIController()
	controllerCtrl := NewControllerController()
	router.HandleFunc("/", apiCtrl.Index).Methods("GET")
	router.HandleFunc("/controllers", controllerCtrl.GetAll).Methods("GET")
	router.HandleFunc("/controllers/{id}", controllerCtrl.Get).Methods("GET")
	router.HandleFunc("/controllers", controllerCtrl.Create).Methods("POST")
	return router
}