/** @jsx React.DOM */

var React = require('react');

// Source types
var GoodNotSource = require('./sources/GoodNot.jsx');


var Source = React.createClass({
    makeSourceDataComponent: function(source) {
        switch( source.type ) {
        case "good":
            return <GoodNotSource />;

        // case "status":
        //     break;

        default:
            throw new Error("unknown source type: " + type);
        };
    },

    render: function() {
        var source = this.props.source,
            dataSection = this.makeSourceDataComponent(source);

        // Each component gets a wrapper with its name and some useful buttons.
        return (
            <div className="panel panel-default">
                <div className="panel-heading">
                    <h3 className="panel-title">
                        {source.name}
                    </h3>
                </div>
                <div className="panel-body">
                    {dataSection}
                </div>
            </div>
        );
    },
});

module.exports = Source;
