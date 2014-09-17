/** @jsx React.DOM */

var React = require('react');


var DashboardContainer = React.createClass({
    render: function() {
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
                </div>
            </div>
        );
    },
});


module.exports = DashboardContainer;
