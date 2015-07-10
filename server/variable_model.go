//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

type Variable struct {
	Model
	Name		string	`json:"name"`
	Type		string	`json:"type"`
	Description	string	`sql:"size:1000" json:"description"`
	RouteID		uint	`sql:"index" json:"routeID,omitempty"`
	ResponseID	uint	`sql:"index" json:"responseID,omitempty"`
}