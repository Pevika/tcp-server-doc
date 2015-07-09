//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

type Response struct {
	Model
	HTTPCode	int
	Description	string		`sql:"size:1000"`
	Variables	[]Variable	`sql:"gorm:many2many:response_variables;"`
	RouteID		int			`sql:"index"`
}

func NewResponse(httpCode int, description string) *Response {
	response := new(Response)
	response.Description = description
	response.HTTPCode = httpCode
	return response
}