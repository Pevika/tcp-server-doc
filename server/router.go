//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	apiCtrl := NewAPIController()
	router.HandleFunc("/", apiCtrl.Index).Methods("GET")
	return router
}