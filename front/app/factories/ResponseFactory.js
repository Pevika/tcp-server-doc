/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
angular.module("app").factory("ResponseFactory", ['ResponseService', '$window', function (ResponseService, $window) {
	
	this.new = function (data) {
		return new $window.Models.Response(data);
	}
	
	this.all = function (routeID) {
		var _this = this;
		return ResponseService.all(routeID).then(function (query) {
			if (query.success) {
				var data = [];
				for (var i = 0 ; i < query.data.responses.length ; ++i) {
					data.push(_this.new(query.data.responses[i]));
				}
				query.data = data;
			}
			return query;
		})
	}
	
	return this;
	
}]);