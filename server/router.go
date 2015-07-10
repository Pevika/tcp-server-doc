//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
)

func Answer(data interface{}, w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	apiCtrl := NewAPIController()
	router.HandleFunc("/", apiCtrl.Index).Methods("GET")
	
	controllerCtrl := NewControllerController()
	router.HandleFunc("/controllers", controllerCtrl.GetAll).Methods("GET")
	router.HandleFunc("/controllers/{id}", controllerCtrl.Get).Methods("GET")
	router.HandleFunc("/controllers", controllerCtrl.Create).Methods("POST")
	router.HandleFunc("/controllers/{id}", controllerCtrl.Update).Methods("PATCH")
	router.HandleFunc("/controllers/{id}", controllerCtrl.Delete).Methods("DELETE")
	
	routeCtrl := NewRouteController()
	router.HandleFunc("/routes/{id}", routeCtrl.Get).Methods("GET")
	router.HandleFunc("/routes", routeCtrl.Create).Methods("POST")
	router.HandleFunc("/routes/{id}", routeCtrl.Update).Methods("PATCH")
	router.HandleFunc("/routes/{id}", routeCtrl.Delete).Methods("DELETE")
	
	router.HandleFunc("/controllers/{cid}/routes", controllerCtrl.GetRoutes).Methods("GET")
	router.HandleFunc("/controllers/{cid}/routes", controllerCtrl.LinkRoute).Methods("POST")
	router.HandleFunc("/controllers/{cid}/routes/{rid}", controllerCtrl.UnlinkRoute).Methods("DELETE")
	
	return router
}