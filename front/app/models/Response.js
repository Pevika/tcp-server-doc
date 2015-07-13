/**
 * @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
 */
 
window.Models.Response = function Response(data) {
	this.id = data.ID;
	this.content = data.content;
	this.description = data.description;
}