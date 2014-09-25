var Flux = require('delorean.js').Flux;

var SourcesApi = require('../api/Sources.js');

var SourceStore = Flux.createStore({
	sources: [],

	actions: {
	},

	initialize: function() {
		// Load the initial state from the server.
		var self = this;

		SourcesApi
			.list()
			.then(function(resp) {
				self.sources = resp.body;
				self.emit('change');
			});
	},

	getState: function() {
		return {
			sources: this.sources,
		};
	},
});

module.exports = SourceStore;
