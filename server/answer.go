//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

type SuccessAnswer struct {
	Success 	bool		`json:"sucess"`
	Data 		interface{}	`json:"data,omitempty"`
}

type RequestError struct {
	Error 	string		`json:"error"`
	Data 	interface{}	`json:"data,omitempty"`
}