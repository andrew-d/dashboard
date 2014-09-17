/** @jsx React.DOM */

var asap = require('asap'),
    Immutable = require('immutable'),
    React = require('react'),
    RRouter = require('rrouter');
var Route = RRouter.Route;

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


var routes = (
    <Route>
        <Route name="index" path="/" view={MainPage} cursor={rootCursor} />
        <Route name="settings" path="/settings" view={MainPage} cursor={rootCursor} />
    </Route>
);

var currentComponent = null;

RRouter.start(routes, function(view) {
    var el = document.querySelector('#application');
    currentComponent = React.renderComponent(view, el);
});
