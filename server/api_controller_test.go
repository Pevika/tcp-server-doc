//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main_test

import (
	"main"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("APIController", func () {
	var request *http.request
	var recorder *httptest.ResponseRecorder
	
	BeforeEach(func() {
		settings := NewSettings()
		settings.SQLDatabase = "test_doc"
		NewServer(settings)
		recorder = httptest.NewRecorder()
	})
	
	AfterEach(func() {
		NewSQL().DB.DropDatabase()
	})
	
	Describe("GET /", func () {
		
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/", nil)
		})
		
		Context("when all is normal", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})
		})
		
	})
	
})