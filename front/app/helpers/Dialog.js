/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */

angular.module("app").service("Dialog", ['$q', function ($q) {
	
	this.confirm = function (message) {
		var deferred = $q.defer();
		var d = confirm(message);
		if (d) {
			deferred.resolve();
		}
		else {
			deferred.reject();
		}
		return deferred.promise;
	}
	
}]);