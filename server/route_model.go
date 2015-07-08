//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"github.com/jinzhu/gorm"
)

type Route struct {
	gorm.Model
	ControllerID 	int			`sql:"index"`
	Name			string
	Description		string		`sql:"size:1000"`
	Responses		[]Response
	Variables		[]Variable
}

func NewRoute(controllerID int, name string, description string) *Route {
	route := new(Route)
	route.ControllerID = controllerID
	route.Name = name
	route.Description = description
	return route
}