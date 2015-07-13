/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
window.Models.Route = function Route(data) {
	this.id = data.ID;
	this.name = data.name;
	this.description = data.description;
	this.route = data.route;
	this.content = data.content;
	this.labelID = data.labelID;
}