//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

type Variable struct {
	Model
	Name		string
	Type		string
	Description	string	`sql:"size:1000"`
}

func NewVariable(name string, type_ string, description string) *Variable {
	variable := new(Variable)
	variable.Name = name
	variable.Type = type_
	variable.Description = description
	return variable
}