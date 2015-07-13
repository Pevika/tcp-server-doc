/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
 angular.module("app").service("RouteService", ['HttpRequest', function (HttpRequest) {
	 
	 this.all = function (controllerID) {
		 return HttpRequest.get("/controllers/" + controllerID + "/routes");
	 }
	 
 }]);