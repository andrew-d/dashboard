/** @jsx React.DOM */

var React = require('react');

var NavBar = require('../components/NavBar.jsx'),
    DashboardContainer = require('../components/DashboardContainer.jsx');


var MainPage = React.createClass({
    render: function() {
        var fullHeight = {
            'height': '100%',
        };

        return (
            <div style={fullHeight}>
                <nav className="navbar navbar-default navbar-fixed-top">
                    <a className="navbar-brand" href="/">Dashboard</a>
                </nav>
                <div className="container-fluid" style={fullHeight}>
                    <div id="layout" className="row" style={fullHeight}>
                        <NavBar />
                        <DashboardContainer />
                    </div>
                </div>
            </div>
        );
    },
});


module.exports = MainPage;
