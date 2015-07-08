//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	settings := NewSettings()
	if settings == nil {
		log.Fatal("Settings are nil")
		return
	}
	router := NewRouter()
	NewSQL()
	log.Println("Server listening on port", settings.Port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(settings.Port), router))
}