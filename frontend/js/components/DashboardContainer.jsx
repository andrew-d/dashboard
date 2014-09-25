/** @jsx React.DOM */

var Flux = require('delorean.js').Flux,
    map = require('lodash-node/compat/collections/map'),
    React = require('react');

var Source = require('./Source.jsx');


var DashboardContainer = React.createClass({
    mixins: [Flux.mixins.storeListener],

    render: function() {
        var sources = this.dispatcher.getStore('sourceStore').sources;

        var sourcesComponents = map(sources, function(v, k) {
            // Each source is half-wide normally, dropping to full-width on tablets.
            return (
                <div className="col-sm-12 col-md-6">
                    <Source source={v} />
                </div>
            );
        });

        return (
            <div id="page-wrapper">
                <div className="container-fluid">
                    <div className="row">
                        <div className="col-lg-12">
                            <h1 className="page-header">
                                Dashboard
                            </h1>
                        </div>
                    </div>

                    <div className="row">
                        {sourcesComponents}
                    </div>
                </div>
            </div>
        );
    },
});


module.exports = DashboardContainer;
