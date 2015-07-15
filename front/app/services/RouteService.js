/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
 angular.module("app").service("RouteService", ['HttpRequest', function (HttpRequest) {
	 
	 this.all = function (controllerID) {
		 return HttpRequest.get("/controllers/" + controllerID + "/routes");
	 }
	 
	 this.create = function (name, description, content) {
		 return HttpRequest.post("/routes", {
			 name: name,
			 description: description,
			 content: content,
			 route: "route"
		 })
	 }
	 
	 this.linkVariable = function (route, variable) {
		 console.log(variable);
		 return HttpRequest.post("/routes/" + route.id + "/variables", {
			 variableID: variable.id
		 })
	 }
	 
 }]);