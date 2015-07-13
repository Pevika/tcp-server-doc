//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type ResponseController struct {
	DB 	*SQL
}

type ResponseSingleAnswer struct {
	Response 	Response 	`json:"Response"`
}

type ResponseCreateRequest struct {
	Description	*string	`json:"description"`
	Content 	*string `json:"content"`
}

func NewResponseController() *ResponseController {
	ctrl := new(ResponseController)
	ctrl.DB = NewSQL()
	return ctrl
}

type ResponseLinkVariableRequest struct {
	VariableID 	*uint	`json:"variableID"`	
}

type ResponseVariablesAnswer struct {
	Variables 	[]Variable	`json:"variables"`
	ResponseID 	uint 		`json:"responseID"`	
}

func (ctrl *ResponseController) Get (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := ResponseSingleAnswer{Response{}}
	if ctrl.DB.DB.Where("ID = ?", id).Find(&data.Response).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)	
	} else {
		Answer(&data, w, 200)
	}
}

func (ctrl *ResponseController) Create (w http.ResponseWriter, r *http.Request) {
	data := ResponseCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else if data.Content == nil {
		Answer(&RequestError{"BadParams", "Content"}, w, 400)
	} else if data.Description == nil {
		Answer(&RequestError{"BadParams", "description"}, w, 400)
	} else {
		Response := Response{Description: *data.Description, Content: *data.Content}
		ctrl.DB.DB.Create(&Response)
		Answer(&ResponseSingleAnswer{Response}, w, 200)	
	}
}

func (ctrl *ResponseController) Update (w http.ResponseWriter, r *http.Request) {
	data := ResponseCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else {
		vars := mux.Vars(r)
		Response := Response{}
		if ctrl.DB.DB.Where("ID = ?", vars["id"]).Find(&Response).RecordNotFound() {
			Answer(&RequestError{"NotFound", nil}, w, 404)
		} else { 
			if data.Description != nil {
				Response.Description = *data.Description
			}
			if data.Content != nil {
				Response.Content = *data.Content
			}
			ctrl.DB.DB.Save(&Response)
			Answer(&ResponseSingleAnswer{Response}, w, 200)
		}
	}
}

func (ctrl *ResponseController) Delete (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Response := Response{}
	if ctrl.DB.DB.Where("ID = ?", vars["id"]).Find(&Response).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)
	} else {
		ctrl.DB.DB.Unscoped().Delete(&Response)
		Answer(&SuccessAnswer{true, nil}, w, 200)
	}
}

func (ctrl *ResponseController) GetVariables(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Response := Response{}
	if ctrl.DB.DB.Where("ID = ?", vars["rid"]).Find(&Response).RecordNotFound() {
		Answer(&RequestError{"NotFound", "Response"}, w, 404)
	} else {
		variables := ResponseVariablesAnswer{[]Variable{}, Response.ID}
		ctrl.DB.DB.Model(&Response).Related(&variables.Variables)
		if variables.Variables == nil {
			variables.Variables = []Variable{}
		}
		Answer(&variables, w, 200)
	}
}

func (ctrl *ResponseController) LinkVariable(w http.ResponseWriter, r *http.Request) {
	data := ResponseLinkVariableRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else if data.VariableID == nil {
		Answer(&RequestError{"BadParams", "variableID"}, w, 400)
	} else {
		vars := mux.Vars(r)
		response := Response{}
		variable := Variable{}
		if ctrl.DB.DB.Where("ID = ?", vars["rid"]).Find(&response).RecordNotFound() {
			Answer(&RequestError{"NotFound", "response"}, w, 404)
		} else if ctrl.DB.DB.Where("ID = ?", *data.VariableID).Find(&variable).RecordNotFound() {
			Answer(&RequestError{"NotFound", "variable"}, w, 404)
		} else if variable.RouteID != 0 || variable.ResponseID != 0 {
			Answer(&RequestError{"VariableAlreadyHasResponse", nil}, w, 400)
		} else {
			variable.ResponseID = response.ID
			ctrl.DB.DB.Save(&variable)
			Answer(&SuccessAnswer{true, nil}, w, 200)	
		}
	}	
}

func (ctrl *ResponseController) UnlinkVariable (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := Response{}
	variable := Variable{}
	if ctrl.DB.DB.Where("ID = ?", vars["rid"]).Find(&response).RecordNotFound() {
		Answer(&RequestError{"NotFound", "response"}, w, 404)
	} else if ctrl.DB.DB.Where("ID = ?", vars["vid"]).Find(&variable).RecordNotFound() {
		Answer(&RequestError{"NotFound", "variable"}, w, 404)
	} else if variable.ResponseID != response.ID {
		Answer(&RequestError{"VariableIsNotInResponse", nil}, w, 400)
	} else {
		variable.ResponseID = 0
		ctrl.DB.DB.Save(&variable)
		Answer(&SuccessAnswer{true, nil}, w, 200)	
	}
}