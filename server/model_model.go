//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"time"
)

type Model struct {
    ID        uint 			`gorm:"primary_key";json:"id"`
    CreatedAt *time.Time	`json:"createdAt,omitempty"`
    UpdatedAt *time.Time	`json:"updatedAt,omitempty"`
    DeletedAt *time.Time	`json:"deletedAt,omitempty"`
}