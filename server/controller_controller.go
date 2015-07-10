//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type ControllerController struct {
	DB	*SQL
}

type ControllerMultipleAnswer struct {
	Controllers	[]Controller	`json:"controllers"`	
}

type ControllerSingleAnswer struct {
	Controller 	Controller 	`json:"controller"`
}

type ControllerCreateRequest struct {
	Name 		*string	`json:"name"`
	Description	*string	`json:"description"`
}

type ControllerRoutesAnswer struct {
	Routes 			[]Route	`json:"routes"`
	ControllerID 	uint	`json:"controllerID"`
}

type ControllerLinkRouteRequest struct {
	RouteID 	*uint	`json:"routeID"`	
}

func NewControllerController() *ControllerController {
	ctrl := new(ControllerController)
	ctrl.DB = NewSQL()
	return ctrl
}

func (ctrl *ControllerController) GetAll (w http.ResponseWriter, r *http.Request) {
	data := ControllerMultipleAnswer{[]Controller{}}
	ctrl.DB.DB.Find(&data.Controllers)
	if data.Controllers == nil {
		data.Controllers = []Controller{}
	}
	Answer(&data, w, 200)
}

func (ctrl *ControllerController) Get (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := ControllerSingleAnswer{Controller{}}
	if ctrl.DB.DB.Where("ID = ?", id).Find(&data.Controller).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)	
	} else {
		Answer(&data, w, 200)
	}
}

func (ctrl *ControllerController) Create (w http.ResponseWriter, r *http.Request) {
	data := ControllerCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else if data.Name == nil || len(*data.Name) == 0 {
		Answer(&RequestError{"BadParams", "name"}, w, 400)
	} else if data.Description == nil {
		Answer(&RequestError{"BadParams", "description"}, w, 400)
	} else {
		controller := Controller{Name: *data.Name, Description: *data.Description}
		ctrl.DB.DB.Create(&controller)
		Answer(&ControllerSingleAnswer{controller}, w, 200)	
	}
}

func (ctrl *ControllerController) Update (w http.ResponseWriter, r *http.Request) {
	data := ControllerCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else {
		vars := mux.Vars(r)
		id := vars["id"]
		controller := Controller{}
		if ctrl.DB.DB.Where("ID = ?", id).Find(&controller).RecordNotFound() {
			Answer(&RequestError{"NotFound", nil}, w, 404)
		} else {
			if data.Name != nil && len(*data.Name) > 0 {
				controller.Name = *data.Name
			}
			if data.Description != nil {
				controller.Description = *data.Description
			}
			ctrl.DB.DB.Save(&controller)
			Answer(&ControllerSingleAnswer{controller}, w, 200)
		}
	}
}

func (ctrl *ControllerController) Delete (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	controller := Controller{}
	if ctrl.DB.DB.Where("ID = ?", vars["id"]).Find(&controller).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)
	} else {
		ctrl.DB.DB.Unscoped().Delete(&controller)
		Answer(&SuccessAnswer{true, nil}, w, 200)
	}
}

func (ctrl *ControllerController) GetRoutes (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	controller := Controller{}
	if ctrl.DB.DB.Where("ID = ?", vars["cid"]).Find(&controller).RecordNotFound() {
		Answer(&RequestError{"NotFound", "controller"}, w, 404)
	} else {
		routes := ControllerRoutesAnswer{[]Route{}, controller.ID}
		ctrl.DB.DB.Model(&controller).Related(&routes.Routes)
		if routes.Routes == nil {
			routes.Routes = []Route{}
		}
		Answer(&routes, w, 200)
	}
}

func (ctrl *ControllerController) LinkRoute (w http.ResponseWriter, r *http.Request) {
	data := ControllerLinkRouteRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else if data.RouteID == nil {
		Answer(&RequestError{"BadParams", "routeID"}, w, 400)
	} else {
		vars := mux.Vars(r)
		controller := Controller{}
		route := Route{}
		if ctrl.DB.DB.Where("ID = ?", vars["cid"]).Find(&controller).RecordNotFound() {
			Answer(&RequestError{"NotFound", "controller"}, w, 404)
		} else if ctrl.DB.DB.Where("ID = ?", *data.RouteID).Find(&route).RecordNotFound() {
			Answer(&RequestError{"NotFound", "route"}, w, 404)
		} else if route.ControllerID != 0 {
			Answer(&RequestError{"RouteAlreadyHasController", nil}, w, 400)
		} else {
			route.ControllerID = controller.ID
			ctrl.DB.DB.Save(&route)
			Answer(&SuccessAnswer{true, nil}, w, 200)	
		}
	}
}

func (ctrl *ControllerController) UnlinkRoute (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	controller := Controller{}
	route := Route{}
	if ctrl.DB.DB.Where("ID = ?", vars["cid"]).Find(&controller).RecordNotFound() {
		Answer(&RequestError{"NotFound", "controller"}, w, 404)
	} else if ctrl.DB.DB.Where("ID = ?", vars["rid"]).Find(&route).RecordNotFound() {
		Answer(&RequestError{"NotFound", "route"}, w, 404)
	} else if route.ControllerID != controller.ID {
		Answer(&RequestError{"RouteIsNotInController", nil}, w, 400)
	} else {
		route.ControllerID = 0
		ctrl.DB.DB.Save(&route)
		Answer(&SuccessAnswer{true, nil}, w, 200)	
	}
}