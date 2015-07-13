/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
window.Models.Variable = function Variable(data) {
	this.id = data.ID;
	this.name = data.name;
	this.type = data.type;
	this.description = data.description;
	this.routeID = data.routeID;
}