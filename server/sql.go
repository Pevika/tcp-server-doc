//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type SQL struct {
	DB	*gorm.DB
}

var _sql *SQL = nil

func NewSQL() *SQL {
	if _sql == nil {
		_sql = new(SQL)
		_sql.Init()
	}
	return _sql
}

func (this *SQL) Init() bool {
	settings := NewSettings()
	db, err := gorm.Open(settings.SQLDriver, settings.SQLUsername + ":" + settings.SQLPassword + "@" + settings.SQLHostname + "/" + settings.SQLDatabase + "?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		this.DB = &db
		return true
	} else {
		log.Fatal("Error while connecting to driver ", settings.SQLDriver, " ", err)
		return false
	}
}