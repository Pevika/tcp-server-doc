/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */

angular.module("app").controller("ControllerCtrl", ['ControllerFactory', '$scope', 'RouteFactory', '$sce', 'VariableFactory', 'ResponseFactory',
	function (ControllerFactory, $scope, RouteFactory, $sce, VariableFactory, ResponseFactory) {

		$scope.controllers = [];

		$scope.init = function () {
			ControllerFactory.all().then(function (query) {
				if (query.success) {
					$scope.controllers = query.data;
					for (var i = 0 ; i < $scope.controllers.length ; ++i) {
						$scope.loadRoutes($scope.controllers[i]);						
					}
				}
				else {
					// TODO : error handler
				}
			});
		}
		
		$scope.showController = function (controller) {
			controller.display = controller.display ? false : true;
		}

		$scope.showRoute = function (route) {
			route.display = route.display ? false : true;
		}
		
		
		$scope.loadRoutes = function (controller) {
			RouteFactory.all(controller.id).then(function (query) {
				if (query.success) {
					controller.routes = query.data;
					for (var i = 0 ; i < controller.routes.length ; ++i) {
						$scope.loadRouteVariables(controller.routes[i]);
						$scope.loadResponses(controller.routes[i]);
					}
					console.log(controller);
				}
				else {
					// TODO: error handler
				}
			})
		}
		
		$scope.loadRouteVariables = function (route) {
			VariableFactory.allOfRoute(route.id).then(function (query) {
				if (query.success) {
					route.variables = query.data;
					route.formatedContent = $sce.trustAsHtml($scope.formatJSON(route.content, route.variables));
				}
			})
		}
		
		$scope.loadResponses = function (route) {
			ResponseFactory.all(route.id).then(function (query) {
				if (query.success) {
					route.responses = query.data;
					for (var i = 0 ; i < route.responses.length ; ++i) {
						$scope.loadResponseVariables(route.responses[i]);
					}
				}	
			});
		}
		
		$scope.loadResponseVariables = function (response) {
			VariableFactory.allOfResponse(response.id).then(function (query) {
				if (query.success) {
					response.variables = query.data;
					response.formatedContent = $sce.trustAsHtml($scope.formatJSON(response.content, response.variables));
				}
			});
		}
		
		$scope.formatJSON = function (content, args) {
			content = content.replace(/\t/g, "   ")
							.replace(/\{/g, "<span class='json-syntax-token'>{</span>")
							.replace(/\}/g, "<span class='json-syntax-token'>}</span>")
							.replace(/\[/g, "<span class='json-syntax-token'>[</span>")
							.replace(/\]/g, "<span class='json-syntax-token'>]</span>")
							.replace(/:/g, "<span class='json-syntax-token'>:</span>")
							.replace(/,/g, "<span class='json-syntax-token'>,</span>");
			var tmp = content.split('\n');
			for (var i = 0 ; i < tmp.length ; ++i) {
				tmp[i] = "<span class='json-line'>" + tmp[i] + "</span>\n";
			}
			content = tmp.join("");
			for (var i = 0 ; args && i < args.length ; ++i) {
				content = content.replace(new RegExp(" "  + args[i].name, 'g'), " <span class='parameter'>" + args[i].name + "</span>");
			}
			return content;
		}
	
	}]);