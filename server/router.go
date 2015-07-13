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

func HandleRequest (fn http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, DELETE, OPTIONS")
		fn(w, r)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	apiCtrl := NewAPIController()
	router.HandleFunc("/", HandleRequest(apiCtrl.Index)).Methods("GET")
	
	controllerCtrl := NewControllerController()
	router.HandleFunc("/controllers", HandleRequest(controllerCtrl.GetAll)).Methods("GET")
	router.HandleFunc("/controllers/{id}", HandleRequest(controllerCtrl.Get)).Methods("GET")
	router.HandleFunc("/controllers", HandleRequest(controllerCtrl.Create)).Methods("POST")
	router.HandleFunc("/controllers/{id}", HandleRequest(controllerCtrl.Update)).Methods("PATCH")
	router.HandleFunc("/controllers/{id}", HandleRequest(controllerCtrl.Delete)).Methods("DELETE")
	
	routeCtrl := NewRouteController()
	router.HandleFunc("/routes/{id}", HandleRequest(routeCtrl.Get)).Methods("GET")
	router.HandleFunc("/routes", HandleRequest(routeCtrl.Create)).Methods("POST")
	router.HandleFunc("/routes/{id}", HandleRequest(routeCtrl.Update)).Methods("PATCH")
	router.HandleFunc("/routes/{id}", HandleRequest(routeCtrl.Delete)).Methods("DELETE")
	
	router.HandleFunc("/controllers/{cid}/routes", HandleRequest(controllerCtrl.GetRoutes)).Methods("GET")
	router.HandleFunc("/controllers/{cid}/routes", HandleRequest(controllerCtrl.LinkRoute)).Methods("POST")
	router.HandleFunc("/controllers/{cid}/routes/{rid}", HandleRequest(controllerCtrl.UnlinkRoute)).Methods("DELETE")
	
	variableCtrl := NewVariableController()
	router.HandleFunc("/variables/{id}", HandleRequest(variableCtrl.Get)).Methods("GET")
	router.HandleFunc("/variables", HandleRequest(variableCtrl.Create)).Methods("POST")
	router.HandleFunc("/variables/{id}", HandleRequest(variableCtrl.Update)).Methods("PATCH")
	router.HandleFunc("/variables/{id}", HandleRequest(variableCtrl.Delete)).Methods("DELETE")
	
	router.HandleFunc("/routes/{rid}/variables", HandleRequest(routeCtrl.GetVariables)).Methods("GET")
	router.HandleFunc("/routes/{rid}/variables", HandleRequest(routeCtrl.LinkVariable)).Methods("POST")
	router.HandleFunc("/routes/{rid}/variables/{vid}", HandleRequest(routeCtrl.UnlinkVariable)).Methods("DELETE")

	responseCtrl := NewResponseController()
	router.HandleFunc("/responses/{id}", HandleRequest(responseCtrl.Get)).Methods("GET")
	router.HandleFunc("/responses", HandleRequest(responseCtrl.Create)).Methods("POST")
	router.HandleFunc("/responses/{id}", HandleRequest(responseCtrl.Update)).Methods("PATCH")
	router.HandleFunc("/responses/{id}", HandleRequest(responseCtrl.Delete)).Methods("DELETE")
	
	router.HandleFunc("/routes/{roid}/responses", HandleRequest(routeCtrl.GetResponses)).Methods("GET")
	router.HandleFunc("/routes/{roid}/responses", HandleRequest(routeCtrl.LinkResponse)).Methods("POST")
	router.HandleFunc("/routes/{roid}/responses/{reid}", HandleRequest(routeCtrl.UnlinkResponse)).Methods("DELETE")
	
	router.HandleFunc("/responses/{rid}/variables", HandleRequest(responseCtrl.GetVariables)).Methods("GET")
	router.HandleFunc("/responses/{rid}/variables", HandleRequest(responseCtrl.LinkVariable)).Methods("POST")
	router.HandleFunc("/responses/{rid}/variables/{vid}", HandleRequest(responseCtrl.UnlinkVariable)).Methods("DELETE")

	labelCtrl := NewLabelController()
	router.HandleFunc("/labels", HandleRequest(labelCtrl.GetAll)).Methods("GET")
	router.HandleFunc("/labels/{id}", HandleRequest(labelCtrl.Get)).Methods("GET")
	router.HandleFunc("/labels", HandleRequest(labelCtrl.Create)).Methods("POST")
	router.HandleFunc("/labels/{id}", HandleRequest(labelCtrl.Update)).Methods("PATCH")
	router.HandleFunc("/labels/{id}", HandleRequest(labelCtrl.Delete)).Methods("DELETE")
	
	router.HandleFunc("/routes/{rid}/label", HandleRequest(routeCtrl.SetLabel)).Methods("POST")
	router.HandleFunc("/routes/{rid}/label", HandleRequest(routeCtrl.UnsetLabel)).Methods("DELETE")

	return router
}