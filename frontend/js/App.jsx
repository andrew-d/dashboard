/** @jsx React.DOM */

var asap = require('asap'),
    Immutable = require('immutable'),
    React = require('react'),
    Router = require('react-router-component');

var AboutPage = require('./pages/AboutPage.jsx'),
    MainPage = require('./pages/MainPage.jsx'),
    SourcesApi = require('./api/Sources.js');

// Imported for side effect.
require('superagent-bluebird-promise');


// This is our state.  All persistent data in the application is stored in this
// object.
var rootState = Immutable.fromJS({
    sources: [],
    data: {},
});


// Our main application.
var App = React.createClass({
    componentDidMount: function() {
        var cursor = this.props.cursor;

        // When the application is loaded, we load all sources.
        SourcesApi.list().then(function(resp) {
            cursor.set('sources', Immutable.fromJS(resp.body));
        });
    },

    render: function() {
        var cursor = this.props.cursor;

        return (
            <Router.Locations hash>
                <Router.Location path="/" handler={MainPage} cursor={cursor} />
                <Router.Location path="/settings" handler={MainPage} cursor={cursor} />
                <Router.Location path="/about" handler={AboutPage} cursor={cursor} />
            </Router.Locations>
        );
    },
});


// This callback gets triggered whenever there are any updates to the
// application state.  We schedule a forced re-render whenever this happens.
var willUpdate = false;
var hasUpdated = function(newData) {
    console.log("Re-rendering application...");

    // 'rootState' always points to the latest set of data that we have.
    rootState = newData;

    // Only re-render once per call from the event loop.
    if( !willUpdate ) {
        willUpdate = true;

        asap(function() {
            willUpdate = false;
            renderApplication();
        });
    }
};

// This function actually renders our application.
var renderApplication = function() {
    var cursor = rootState.cursor(hasUpdated);

    React.renderComponent(
        App({cursor: cursor}),
        document.querySelector('#application')
    );
};

// Finally, perform the initial render.
renderApplication();
