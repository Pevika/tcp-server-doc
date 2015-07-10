//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
	"log"
)

func main() {
	settings := NewSettings()
	if settings == nil {
		log.Fatal("Settings are nil")
		return
	}
	NewServer(settings)
}