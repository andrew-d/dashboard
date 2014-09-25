/** @jsx React.DOM */

var React = require('react');

var NavBar = require('../components/NavBar.jsx'),
    DashboardContainer = require('../components/DashboardContainer.jsx');


var MainPage = React.createClass({
    render: function() {
        return (
            <div id='wrapper'>
                <NavBar />
                <DashboardContainer />
            </div>
        );
    },
});


module.exports = MainPage;
