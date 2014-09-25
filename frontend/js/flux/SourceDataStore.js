var Flux = require('delorean.js').Flux;

var SourcesApi = require('../api/Sources.js');

var SourceDataStore = Flux.createStore({
	sourceData: {},

	actions: {
	},

	getState: function() {
		return {
			sourceData: this.sourceData,
		};
	},
});

module.exports = SourceDataStore;
