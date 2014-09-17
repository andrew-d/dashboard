/** @jsx React.DOM */

var React = require('react');

var NavBar = require('../components/NavBar.jsx'),
    DashboardContainer = require('../components/DashboardContainer.jsx');


var MainPage = React.createClass({
    // Debugging
    componentDidMount: function() {
        console.log("cursor: " + this.props.cursor);
    },

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
