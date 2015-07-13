//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

type Controller struct {
	Model
	Name 			string		`json:"name"`
	Description		string		`json:"description"`
}