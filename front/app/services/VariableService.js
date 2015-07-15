/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
 angular.module("app").service("VariableService", ['HttpRequest', function (HttpRequest) {
	 
	 this.allOfRoute = function (routeID) {
		 return HttpRequest.get("/routes/" + routeID + "/variables");
	 }
	 
	 this.allOfResponse = function (responseID) {
		 return HttpRequest.get("/responses/" + responseID + "/variables");
	 }
	 
	 this.create = function (name, type, description) {
		 return HttpRequest.post("/variables", {
			 name: name,
			 type: type,
			 description: description
		 })
	 }
	 
 }]);