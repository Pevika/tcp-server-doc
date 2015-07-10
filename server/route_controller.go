//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
)

type RouteController struct {
	DB	*SQL
}

type RouteSingleAnswer struct {
	Route 	Route 	`json:"route"`
}

type RouteCreateRequest struct {
	Name 			*string	`json:"name"`
	Description 	*string	`json:"description"`
	Route 			*string `json:"route"`
	Content			*string `json:"content"`
}

func NewRouteController() *RouteController {
	ctrl := new(RouteController)
	ctrl.DB = NewSQL()
	return ctrl
}

func (ctrl *RouteController) Get (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := RouteSingleAnswer{Route{}}
	if ctrl.DB.DB.Where("ID = ?", vars["id"]).Find(&data.Route).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)	
	} else {
		Answer(&data, w, 200)
	}
}

func (ctrl *RouteController) Create (w http.ResponseWriter, r *http.Request) {
	data := RouteCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
		log.Printf("%v\n", err)		
	} else if data.Name == nil || len(*data.Name) == 0 {
		Answer(&RequestError{"BadParams", "name"}, w, 400)
	} else if data.Description == nil {
		Answer(&RequestError{"BadParams", "description"}, w, 400)
	} else if data.Route == nil || len(*data.Route) == 0 {
		Answer(&RequestError{"BadParams", "route"}, w, 400)
	} else if data.Content == nil || len(*data.Content) == 0 {
		Answer(&RequestError{"BadParams", "route"}, w, 400)
	} else {
		route := Route{Name: *data.Name, Description: *data.Description, Route: *data.Route, Content: *data.Content}
		ctrl.DB.DB.Create(&route)
		Answer(&RouteSingleAnswer{route}, w, 200)	
	}
}

func (ctrl *RouteController) Update (w http.ResponseWriter, r *http.Request) {
	data := RouteCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else {
		vars := mux.Vars(r)
		route := Route{}
		if ctrl.DB.DB.Where("ID = ?", vars["id"]).Find(&route).RecordNotFound() {
			Answer(&RequestError{"NotFound", nil}, w, 404)
		} else { 
			if data.Name != nil && len(*data.Name) > 0 {
				route.Name = *data.Name
			}
			if data.Description != nil {
				route.Description = *data.Description
			}
			if data.Content != nil && len(*data.Content) > 0 {
				route.Content = *data.Content
			}
			if data.Route != nil && len(*data.Route) > 0 {
				route.Route = *data.Route
			}
			ctrl.DB.DB.Save(&route)
			Answer(&RouteSingleAnswer{route}, w, 200)
		}
	}
}

func (ctrl *RouteController) Delete (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route := Route{}
	if ctrl.DB.DB.Where("ID = ?", vars["id"]).Find(&route).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)
	} else {
		ctrl.DB.DB.Unscoped().Delete(&route)
		Answer(true, w, 200)
	}
}