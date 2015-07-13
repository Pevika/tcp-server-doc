/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
angular.module("app").factory("QueryStatus", ['$location', function ($location) {

	function QueryStatus(queryResult, success) {
		this.status = queryResult.status;
		this.data = queryResult.data;
		if (this.data && this.data.error) {
			this.error = this.data.error;
		}
		this.success = success;
		if (this.status == 0 && navigator.onLine == false) {
			this.status = 410;
		}
	};

	QueryStatus.prototype.getErrorMessage = function () {
		var message;
		if (this.error) {
			message = "";
			if (this.error instanceof Array) {
				message += this.error.join(". ");
			}
			else {
				message += this.error;				
			}
		}
		else {
		 	message = "Erreur " + this.status;
		}
		return message;
	};

	return QueryStatus;

}]);