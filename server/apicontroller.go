package main

import (
	"net/http"
	"fmt"
	"encoding/json"
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
	fmt.Printf("%+v\n", data)
	json.NewEncoder(w).Encode(data)
}