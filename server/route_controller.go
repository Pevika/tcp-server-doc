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

type RouteVariablesAnswer struct {
	Variables 	[]Variable 	`json:"variables"`
	RouteID		uint		`json:"routeID"`
}

type RouteLinkVariableRequest struct {
	VariableID 	*uint	`json:"variableID"`	
}

type RouteLinkResponseRequest struct {
	ResponseID	*uint	`json:"responseID"`
}

type RouteResponsesAnswer struct {
	Responses  	[]Response	`json:"responses"`
	RouteID 	uint		`json:"routeID"`
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
		Answer(&SuccessAnswer{true, nil}, w, 200)
	}
}

func (ctrl *RouteController) GetVariables(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route := Route{}
	if ctrl.DB.DB.Where("ID = ?", vars["rid"]).Find(&route).RecordNotFound() {
		Answer(&RequestError{"NotFound", "route"}, w, 404)
	} else {
		variables := RouteVariablesAnswer{[]Variable{}, route.ID}
		ctrl.DB.DB.Model(&route).Related(&variables.Variables)
		if variables.Variables == nil {
			variables.Variables = []Variable{}
		}
		Answer(&variables, w, 200)
	}
}

func (ctrl *RouteController) LinkVariable(w http.ResponseWriter, r *http.Request) {
	data := RouteLinkVariableRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
		log.Printf("%v\n", err)		
	} else if data.VariableID == nil {
		Answer(&RequestError{"BadParams", "variableID"}, w, 400)
	} else {
		vars := mux.Vars(r)
		route := Route{}
		variable := Variable{}
		if ctrl.DB.DB.Where("ID = ?", vars["rid"]).Find(&route).RecordNotFound() {
			Answer(&RequestError{"NotFound", "route"}, w, 404)
		} else if ctrl.DB.DB.Where("ID = ?", *data.VariableID).Find(&variable).RecordNotFound() {
			Answer(&RequestError{"NotFound", "variable"}, w, 404)
		} else if variable.RouteID != 0 || variable.ResponseID != 0 {
			Answer(&RequestError{"VariableAlreadyHasRoute", nil}, w, 400)
		} else {
			variable.RouteID = route.ID
			ctrl.DB.DB.Save(&variable)
			Answer(&SuccessAnswer{true, nil}, w, 200)	
		}
	}	
}

func (ctrl *RouteController) UnlinkVariable (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route := Route{}
	variable := Variable{}
	if ctrl.DB.DB.Where("ID = ?", vars["rid"]).Find(&route).RecordNotFound() {
		Answer(&RequestError{"NotFound", "route"}, w, 404)
	} else if ctrl.DB.DB.Where("ID = ?", vars["vid"]).Find(&variable).RecordNotFound() {
		Answer(&RequestError{"NotFound", "variable"}, w, 404)
	} else if variable.RouteID != route.ID {
		Answer(&RequestError{"VariableIsNotInRoute", nil}, w, 400)
	} else {
		variable.RouteID = 0
		ctrl.DB.DB.Save(&variable)
		Answer(&SuccessAnswer{true, nil}, w, 200)	
	}
}

func (ctrl *RouteController) GetResponses (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route := Route{}
	if ctrl.DB.DB.Where("ID = ?", vars["roid"]).Find(&route).RecordNotFound() {
		Answer(&RequestError{"NotFound", "route"}, w, 404)
	} else {
		responses := RouteResponsesAnswer{[]Response{}, route.ID}
		ctrl.DB.DB.Model(&route).Related(&responses.Responses)
		if responses.Responses == nil {
			responses.Responses = []Response{}
		}
		Answer(&responses, w, 200)
	}
}

func (ctrl *RouteController) LinkResponse (w http.ResponseWriter, r *http.Request) {
	data := RouteLinkResponseRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else if data.ResponseID == nil {
		Answer(&RequestError{"BadParams", "responseID"}, w, 400)
	} else {
		vars := mux.Vars(r)
		route := Route{} 
		response := Response{}
		if ctrl.DB.DB.Where("ID = ?", vars["roid"]).Find(&route).RecordNotFound() {
			Answer(&RequestError{"NotFound", "route"}, w, 404)
		} else if ctrl.DB.DB.Where("ID = ?", *data.ResponseID).Find(&response).RecordNotFound() {
			Answer(&RequestError{"NotFound", "response"}, w, 404)
		} else if response.RouteID != 0 {
			Answer(&RequestError{"ResponseAlreadyHasRoute", nil}, w, 400)
		} else {
			response.RouteID = route.ID
			ctrl.DB.DB.Save(&response)
			Answer(&SuccessAnswer{true, nil}, w, 200)
		}
	}
}

func (ctrl *RouteController) UnlinkResponse (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	route := Route{}
	response := Response{}
	if ctrl.DB.DB.Where("ID = ?", vars["roid"]).Find(&route).RecordNotFound() {
			Answer(&RequestError{"NotFound", "route"}, w, 404)
	} else if ctrl.DB.DB.Where("ID = ?", vars["reid"]).Find(&response).RecordNotFound() {
		Answer(&RequestError{"NotFound", "response"}, w, 404)
	} else if response.RouteID != route.ID {
		Answer(&RequestError{"ResponseIsNotInRoute", nil}, w, 400)
	} else {
		response.RouteID = 0
		ctrl.DB.DB.Save(&response)
		Answer(&SuccessAnswer{true, nil}, w, 200)	
	}
}