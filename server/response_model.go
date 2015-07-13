//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

type Response struct {
	Model
	Content		string		`sql:"size:1000" json:"content"`
	Description	string		`sql:"size:1000" json:"description"`
	RouteID		uint		`sql:"index" json:"routeID,omitempty"`
}