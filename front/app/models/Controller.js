/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
window.Models.Controller = function Controller(data) {
	this.hydrate(data);
}

window.Models.Controller.prototype.hydrate = function (data) {
	this.id = data.ID;
	this.name = data.name;
	this.description = data.description;	
}