//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type LabelController struct {
	DB	*SQL
}

type LabelMultipleAnswer struct {
	Labels	[]Label	`json:"labels"`	
}

type LabelSingleAnswer struct {
	Label 	Label 	`json:"label"`
}

type LabelCreateRequest struct {
	Name 		*string	`json:"name"`
	Color		*string	`json:"color"`
}

func NewLabelController() *LabelController {
	ctrl := new(LabelController)
	ctrl.DB = NewSQL()
	return ctrl
}

func (ctrl *LabelController) GetAll (w http.ResponseWriter, r *http.Request) {
	data := LabelMultipleAnswer{[]Label{}}
	ctrl.DB.DB.Find(&data.Labels)
	if data.Labels == nil {
		data.Labels = []Label{}
	}
	Answer(&data, w, 200)
}

func (ctrl *LabelController) Get (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := LabelSingleAnswer{Label{}}
	if ctrl.DB.DB.Where("ID = ?", id).Find(&data.Label).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)	
	} else {
		Answer(&data, w, 200)
	}
}

func (ctrl *LabelController) Create (w http.ResponseWriter, r *http.Request) {
	data := LabelCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else if data.Name == nil || len(*data.Name) == 0 {
		Answer(&RequestError{"BadParams", "name"}, w, 400)
	} else if data.Color == nil || len(*data.Color) == 0 {
		Answer(&RequestError{"BadParams", "color"}, w, 400)
	} else {
		Label := Label{Name: *data.Name, Color: *data.Color}
		ctrl.DB.DB.Create(&Label)
		Answer(&LabelSingleAnswer{Label}, w, 200)	
	}
}

func (ctrl *LabelController) Update (w http.ResponseWriter, r *http.Request) {
	data := LabelCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		Answer(&RequestError{"BadParams", nil}, w, 400)
	} else {
		vars := mux.Vars(r)
		id := vars["id"]
		Label := Label{}
		if ctrl.DB.DB.Where("ID = ?", id).Find(&Label).RecordNotFound() {
			Answer(&RequestError{"NotFound", nil}, w, 404)
		} else {
			if data.Name != nil && len(*data.Name) > 0 {
				Label.Name = *data.Name
			}
			if data.Color != nil && len(*data.Color) > 0 {
				Label.Color = *data.Color
			}
			ctrl.DB.DB.Save(&Label)
			Answer(&LabelSingleAnswer{Label}, w, 200)
		}
	}
}

func (ctrl *LabelController) Delete (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Label := Label{}
	if ctrl.DB.DB.Where("ID = ?", vars["id"]).Find(&Label).RecordNotFound() {
		Answer(&RequestError{"NotFound", nil}, w, 404)
	} else {
		ctrl.DB.DB.Unscoped().Delete(&Label)
		Answer(&SuccessAnswer{true, nil}, w, 200)
	}
}