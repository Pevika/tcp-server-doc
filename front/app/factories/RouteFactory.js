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
	
	this.create = function (name, description, content) {
		var _this = this;
		return RouteService.create(name, description, content).then(function (query) {
			if (query.success) {
				query.data = _this.new(query.data.route);
			}
			return query;
		})
	}
	
	this.linkVariable = function (route, variable) {
		return RouteService.linkVariable(route, variable).then(function (query) {
			if (query.success) {
				if (!route.variables) {
					route.variables = [];
				}
				route.variables.push(variable);
			}
			return query;
		})
	}
	
	return this;
	
}]);