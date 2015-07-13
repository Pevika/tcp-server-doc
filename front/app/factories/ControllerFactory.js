/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
angular.module("app").factory("ControllerFactory", ['ControllerService', '$window',
	function (ControllerService, $window) {
	
		this.new = function (data) {
			return new $window.Models.Controller(data);
		}
	
		this.all = function () {
			var _this = this;
			return ControllerService.all().then(function (query) {
				if (query.success) {
					var data = [];
					console.log(query.data);
					for (var i = 0 ; i < query.data.controllers.length ; ++i) {
						data.push(_this.new(query.data.controllers[i]));
					}
					query.data = data;
				}
				return query;
			});
		}
		
		this.loadRoutes = function (id) {
			
		}
	
		return this;
	
	}]);