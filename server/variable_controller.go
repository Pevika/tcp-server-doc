//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type VariableController struct {
	DB 	*SQL
}

type VariableSingleAnswer struct {
	Variable 	Variable 	`json:"variable"`
}

type VariableCreateRequest struct {
	Name 		*string	`json:"name"`
	Type		*string	`json:"type"`
	Description	*string	`json:"description"`
}

func NewVariableController() *VariableController {
	ctrl := new(VariableController)
	ctrl.DB = NewSQL()
	return ctrl
}

func (ctrl *VariableController) Get (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := VariableSingleAnswer{Variable{}}
	if ctrl.DB.DB.Where("ID = ?", id).Find(&data.Variable).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)	
	} else {
		Answer(&data, w, 200)
	}
}

func (ctrl *VariableController) Create (w http.ResponseWriter, r *http.Request) {
	data := VariableCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else if data.Name == nil || len(*data.Name) == 0 {
		Answer(&RequestError{"BadParams", "name"}, w, 400)
	} else if data.Description == nil {
		Answer(&RequestError{"BadParams", "description"}, w, 400)
	} else if data.Type == nil || len(*data.Type) == 0 {
		Answer(&RequestError{"BadParams", "type"}, w, 400)
	} else {
		variable := Variable{Name: *data.Name, Description: *data.Description, Type: *data.Type}
		ctrl.DB.DB.Create(&variable)
		Answer(&VariableSingleAnswer{variable}, w, 200)	
	}
}

func (ctrl *VariableController) Update (w http.ResponseWriter, r *http.Request) {
	data := VariableCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else {
		vars := mux.Vars(r)
		variable := Variable{}
		if ctrl.DB.DB.Where("ID = ?", vars["id"]).Find(&variable).RecordNotFound() {
			Answer(&RequestError{"NotFound", nil}, w, 404)
		} else { 
			if data.Name != nil && len(*data.Name) > 0 {
				variable.Name = *data.Name
			}
			if data.Description != nil {
				variable.Description = *data.Description
			}
			if data.Type != nil && len(*data.Type) > 0 {
				variable.Type = *data.Type
			}
			ctrl.DB.DB.Save(&variable)
			Answer(&VariableSingleAnswer{variable}, w, 200)
		}
	}
}

func (ctrl *VariableController) Delete (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	variable := Variable{}
	if ctrl.DB.DB.Where("ID = ?", vars["id"]).Find(&variable).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)
	} else {
		ctrl.DB.DB.Unscoped().Delete(&variable)
		Answer(&SuccessAnswer{true, nil}, w, 200)
	}
}