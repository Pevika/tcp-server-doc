/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */

angular.module("app").service("ResponseService", ['HttpRequest', function (HttpRequest) {
	
	this.all = function (routeID) {
		return HttpRequest.get("/routes/" + routeID + "/responses");
	}
	
}]);