//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"github.com/jinzhu/gorm"
)

type Controller struct {
	gorm.Model
	Description		string
	Routes 			[]Route
}

func NewController(description string) *Controller {
	controller := new(Controller)
	controller.Description = description
	return controller
}