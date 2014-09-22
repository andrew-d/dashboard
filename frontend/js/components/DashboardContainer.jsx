/** @jsx React.DOM */

var React = require('react'),
    map = require('lodash-node/compat/collections/map');

var GoodNotSource = require('./sources/GoodNot.jsx');


var DashboardContainer = React.createClass({
    render: function() {
        var sources = this.props.cursor.deref();
        var sourcesComponents = sources.map(function(v, k) {
            var cursor = this.props.cursor.cursor(k),
                type = v.get('type');

            switch( type ) {
            case "good":
                return <GoodNotSource cursor={cursor} />;

            // case "status":
            //     break;

            default:
                throw new Error("unknown source type: " + type);
            };
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
                        <div className="col-lg-12">
                            {sourcesComponents}
                        </div>
                    </div>
                </div>
            </div>
        );
    },
});


module.exports = DashboardContainer;
