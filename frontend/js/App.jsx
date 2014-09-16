/** @jsx React.DOM */

var React = require('react'),
    RRouter = require('rrouter');
var Route = RRouter.Route;

var MainPage = require('./pages/MainPage.jsx');


// TODO:
//  - Set up immutable cursors for app state
//  - Load Socket.IO and listen for changes
//  - Pass state to current route


var routes = (
    <Route>
        <Route path="/" view={MainPage} />
    </Route>
);

RRouter.start(routes, function(view) {
    var el = document.querySelector('#application');
    React.renderComponent(view, el);
});
