/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail;com>
 */
 
angular.module("app").factory("VariableFactory", ['VariableService', '$window', function (VariableService, $window) {
	
	this.new = function (data) {
		return new $window.Models.Variable(data);
	}
	
	this.allOfRoute = function (routeID) {
		var _this = this;
		return VariableService.allOfRoute(routeID).then(function (query) {
			if (query.success) {
				var data = [];
				for (var i = 0 ; i < query.data.variables.length ; ++i) {
					data.push(_this.new(query.data.variables[i]));
				}
				query.data = data;
			}
			return query;
		});
	}
	
	this.allOfResponse = function (responseID) {
		var _this = this;
		return VariableService.allOfResponse(responseID).then(function (query) {
			if (query.success) {
				var data = [];
				for (var i = 0 ; i < query.data.variables.length ; ++i) {
					data.push(_this.new(query.data.variables[i]));
				}
				query.data = data;
			}
			return query;
		});
	}
	
	return this;
	
}]);