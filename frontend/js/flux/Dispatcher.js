var Flux = require('delorean.js').Flux;

var SourceStore = require('./SourceStore.js');


var Dispatcher = Flux.createDispatcher({
	getStores: function() {
		return {
			sourceStore: new SourceStore(),
		};
	},
});

module.exports = Dispatcher;
