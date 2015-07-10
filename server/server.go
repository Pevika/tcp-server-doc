//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"log"
	"net/http"
	"strconv"
)

func NewServer(settings *Settings) {
	router := NewRouter()
	NewSQL()
	log.Println("Server listening on port", settings.Port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(settings.Port), router))
}