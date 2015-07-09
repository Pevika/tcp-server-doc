//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"net/http"
)

type APIController struct {
	Data *API
}

func NewAPIController() *APIController {
	ctrl := new(APIController)
	ctrl.Data = NewAPI()
	return ctrl
}

func (ctrl *APIController) Index (w http.ResponseWriter, r *http.Request) {
	data := struct {
		Api	API	`json:"api"`
	}{
		*ctrl.Data,
	}
	Answer(&data, w, 200)
}