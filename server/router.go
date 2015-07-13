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
	
	variableCtrl := NewVariableController()
	router.HandleFunc("/variables/{id}", variableCtrl.Get).Methods("GET")
	router.HandleFunc("/variables", variableCtrl.Create).Methods("POST")
	router.HandleFunc("/variables/{id}", variableCtrl.Update).Methods("PATCH")
	router.HandleFunc("/variables/{id}", variableCtrl.Delete).Methods("DELETE")
	
	router.HandleFunc("/routes/{rid}/variables", routeCtrl.GetVariables).Methods("GET")
	router.HandleFunc("/routes/{rid}/variables", routeCtrl.LinkVariable).Methods("POST")
	router.HandleFunc("/routes/{rid}/variables/{vid}", routeCtrl.UnlinkVariable).Methods("DELETE")

	responseCtrl := NewResponseController()
	router.HandleFunc("/responses/{id}", responseCtrl.Get).Methods("GET")
	router.HandleFunc("/responses", responseCtrl.Create).Methods("POST")
	router.HandleFunc("/responses/{id}", responseCtrl.Update).Methods("PATCH")
	router.HandleFunc("/responses/{id}", responseCtrl.Delete).Methods("DELETE")
	
	router.HandleFunc("/routes/{roid}/responses", routeCtrl.GetResponses).Methods("GET")
	router.HandleFunc("/routes/{roid}/responses", routeCtrl.LinkResponse).Methods("POST")
	router.HandleFunc("/routes/{roid}/responses/{reid}", routeCtrl.UnlinkResponse).Methods("DELETE")
	
	router.HandleFunc("/responses/{rid}/variables", responseCtrl.GetVariables).Methods("GET")
	router.HandleFunc("/responses/{rid}/variables", responseCtrl.LinkVariable).Methods("POST")
	router.HandleFunc("/responses/{rid}/variables/{vid}", responseCtrl.UnlinkVariable).Methods("DELETE")

	return router
}