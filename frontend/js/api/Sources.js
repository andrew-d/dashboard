var isNumber = require('lodash-node/compat/objects/isNumber'),
	request = require('superagent');


var SourcesApi = {
	list: function() {
		return request
			.get('/api/sources')
			.set('Accept', 'application/json')
			.promise();
	},

	get: function(id) {
		if( !isNumber(id) ) {
			throw new Error("'id' must be a number");
		}

		return request
			.get('/api/sources/' + id)
			.set('Accept', 'application/json')
			.promise();
	},

	add: function(name, type, settings) {
		if( !name || !type ) {
			throw new Error("'name' and 'type' are required");
		}
		if( !settings ) {
			settings = {};
		}

		return request
			.post('/api/sources')
			.send({name: name, type: type, settings: settings})
			.set('Accept', 'application/json')
			.promise();
	},

	'delete': function(id) {
		if( !isNumber(id) ) {
			throw new Error("'id' must be a number");
		}

		return request
			.delete('/api/sources/' + id)
			.set('Accept', 'application/json')
			.promise();
	},

	getData: function(id) {
		return request
			.get('/api/sources/' + id + '/data')
			.set('Accept', 'application/json')
			.promise();
	},
};

module.exports = SourcesApi;
