//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"os"
	"encoding/json"
	"log"
)

type Settings struct {
	SQLDriver		string	`json:"SQL_driver"`
	SQLHostname		string	`json:"SQL_hostname"`
	SQLUsername		string	`json:"SQL_username"`
	SQLPassword		string	`json:"SQL_password"`
	SQLDatabase		string	`json:"SQL_database"`
	SQLTablePrefix	string	`json:"SQL_table_prefix"`
	Port			int		`json:"port"`
}

var settings *Settings = nil

func NewSettings() *Settings {
	if settings == nil {
		settings = new(Settings)
		if settings.Init() == false {
			return nil
		}	
	}
	return settings
}

func (this *Settings) Init() bool {
	configFile, err := os.Open("config.json")
	if err == nil {
		parser := json.NewDecoder(configFile)
		err = parser.Decode(&this)
		if err == nil {
			return true
		} else {
			log.Fatal("Error while loading settings file ", err.Error())
		}
	} else {
		log.Fatal("Error while loading settings file ", err.Error())
	}
	return false
}