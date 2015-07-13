/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */

angular.module("app").service("ControllerService", ['HttpRequest', function (HttpRequest) {
	
	this.all = function () {
		return HttpRequest.get("/controllers");
	}
	
}]);