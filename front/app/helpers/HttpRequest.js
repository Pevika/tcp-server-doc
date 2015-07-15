/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */

angular.module("app").factory("HttpRequest", ['$http', 'QueryStatus',
	function ($http, QueryStatus) {

		var API_URL = "http://localhost:4444";

		this.get = function (url) {
			return $http({
				method: "GET",
				url: API_URL + url,
				headers: {
					'Content-Type': "application/json"
				}
			}).then(function (result) {
				return new QueryStatus(result, true);
			}, function (result) {
				var query = new QueryStatus(result, false);
				return query;
			})
		}

		this.post = function (url, data) {
			return $http({
				method: "POST",
				url: API_URL + url,
				headers: {
					'Content-Type': "application/json"
				},
				data: data
			}).then(function (result) {
				return new QueryStatus(result, true);
			}, function (result) {
				return new QueryStatus(result, false);
			})
		}

		this.put = function (url, data) {
			return $http({
				method: "PUT",
				url: API_URL + url,
				headers: {
					'Content-Type': "application/json"
				},
				data: data
			}).then(function (result) {
				return new QueryStatus(result, true);
			}, function (result) {
				return new QueryStatus(result, false);
			})
		}
		
		this.delete = function (url) {
			return $http({
				method: "DELETE",
				url: API_URL + url,
				headers: {
					'Content-Type': "application/json"
				}
			}).then(function (result) {
				return new QueryStatus(result, true);
			}, function (result) {
				var query = new QueryStatus(result, false);
				return query;
			})
		}
		
		this.patch = function (url, data) {
			return $http({
				method: "PATCH",
				url: API_URL + url,
				headers: {
					'Content-Type': "application/json"
				},
				data: data
			}).then(function (result) {
				return new QueryStatus(result, true);
			}, function (result) {
				return new QueryStatus(result, false);
			})
		}

		return this;

	}]);	