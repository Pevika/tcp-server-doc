package main

import (
	"time"
)

type API struct {
	Version	string		`json:"version"`
	Date	time.Time	`json:"date"`
}

func NewAPI() *API {
	api := new(API)
	api.Version = "0.0.0"
	api.Date = time.Now()
	return api
}