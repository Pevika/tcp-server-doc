/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
angular.module("app").factory("RouteFactory", ['RouteService', '$window',
	function (RouteService, $window) {
	
	this.new = function (data) {
		return new $window.Models.Route(data);
	}
	
	this.all = function (controllerID) {
		var _this = this;
		return RouteService.all(controllerID).then(function (query) {
			if (query.success) {
				var data = [];
				for (var i = 0 ; i < query.data.routes.length ; ++i) {
					data.push(_this.new(query.data.routes[i]));
				}
				query.data = data;
			}
			return query;
		});
	}
	
	return this;
	
}]);