/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */

angular.module("app").service("ControllerService", ['HttpRequest', function (HttpRequest) {
	
	this.all = function () {
		return HttpRequest.get("/controllers");
	}
	
	this.create = function (name, description) {
		return HttpRequest.post("/controllers", {
			name: name,
			description: description
		});
	}
	
	this.update = function (controller, name, description) {
		return HttpRequest.patch("/controllers/" + controller.id, {
			name: name,
			description: description
		})
	}
	
	this.delete = function (controller) {
		return HttpRequest.delete("/controllers/" + controller.id);
	}
	
	this.linkRoute = function (controller, route) {
		return HttpRequest.post("/controllers/" + controller.id + "/routes", {
			routeID: route.id
		});
	}
	
}]);