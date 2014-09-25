/** @jsx React.DOM */

var Flux = require('delorean.js').Flux,
    Promise = require('bluebird'),
    React = require('react'),
    Router = require('react-router-component');

// Imported for side effect.
require('superagent-bluebird-promise');

// Tell Delorean.js to use Bluebird promises instead of its default,
// in order to save on minified space.
Flux.define('Promise', Promise);

var Dispatcher = require('./flux/Dispatcher.js');

var AboutPage = require('./pages/AboutPage.jsx'),
    MainPage = require('./pages/MainPage.jsx');


// Our main application.
var App = React.createClass({
    mixins: [Flux.mixins.storeListener],

    render: function() {
        return (
            <Router.Locations hash>
                <Router.Location path="/" handler={MainPage} />
                <Router.Location path="/settings" handler={MainPage} />
                <Router.Location path="/about" handler={AboutPage} />
            </Router.Locations>
        );
    },
});

React.renderComponent(
    App({dispatcher: Dispatcher}),
    document.querySelector('#application')
);
