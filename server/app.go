package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	
	apiCtrl := NewAPIController()
	
	router.HandleFunc("/", apiCtrl.Index).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":4444", router))
}