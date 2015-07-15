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
		
		this.create = function (name, description) {
			var _this = this;
			return ControllerService.create(name, description).then(function (query) {
				if (query.success) {
					query.data = _this.new(query.data.controller);
				}
				return query;
			});
		}
		
		this.update = function (controller, name, description) {
			return ControllerService.update(controller, name, description).then(function (query) {
				if (query.success) {
					controller.hydrate(query.data.controller);					
				}
				return query;
			})
		}
		
		this.delete = function (controller) {
			return ControllerService.delete(controller);
		}
	
		return this;
	
	}]);