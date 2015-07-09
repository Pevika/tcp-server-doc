//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

type Controller struct {
	Model
	Name 			string		`json:"name"`
	Description		string		`json:"description"`
	Routes 			[]Route		`json:"routes"`
}

func NewController(description string) *Controller {
	controller := new(Controller)
	controller.Description = description
	return controller
}