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
	/*c := cors.New(cors.Options{
		AllowOriginFunc: func (origin string) bool {
			log.Printf("Origin %s is allowed", origin)
			return true
		},
		AllowedHeaders: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: false,
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	})
	handler := c.Handler(router)*/
	log.Println("Server listening on port", settings.Port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(settings.Port), router))
}