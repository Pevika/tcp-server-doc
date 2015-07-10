//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main_test

import (
	"main"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpects(t, "Server Suite")
}