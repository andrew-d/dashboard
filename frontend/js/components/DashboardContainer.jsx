/** @jsx React.DOM */

var React = require('react'),
    map = require('lodash-node/compat/collections/map');

var Source = require('./Source.jsx');


var DashboardContainer = React.createClass({
    render: function() {
        var sources = this.props.cursor.deref();
        var sourcesComponents = sources.map(function(v, k) {
			var cursor = this.props.cursor.cursor(k);

			// Each source is half-wide normally, dropping to full-width on tablets.
			return (
				<div className="col-sm-12 col-md-6">
					<Source cursor={cursor} />
				</div>
			);
        }.bind(this)).toArray();

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
