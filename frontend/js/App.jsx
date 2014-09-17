/** @jsx React.DOM */

var asap = require('asap'),
    Immutable = require('immutable'),
    React = require('react'),
    Router = require('react-router-component');

var MainPage = require('./pages/MainPage.jsx');


// TODO:
//  - Set up immutable cursors for app state
//  - Load Socket.IO and listen for changes
//  - Pass state to current route

var state = Immutable.fromJS({
    dashboards: [],
    data: {},
});

var rootCursor = state.cursor(function() {
    if( currentComponent !== null ) {
        console.log("Re-rendering application...");
        currentComponent.forceUpdate();
    }
});


// TODO:
//  - load data into immutable state
//  - pass state to page
//  - have page render state
//  - when the page is changed, force a re-render

var App = React.createClass({
    render: function() {
        return (
            <Router.Locations>
                <Router.Location path="/" handler={MainPage} cursor={rootCursor} />
                <Router.Location path="/settings" handler={MainPage} cursor={rootCursor} />
            </Router.Locations>
        );
    },
});


React.renderComponent(App(), document.querySelector('#application'));
