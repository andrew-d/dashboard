/** @jsx React.DOM */

var React = require('react'),
    Router = require('react-router-component'),
    map = require('lodash-node/compat/collections/map');


// Navbar link that applies the 'active' class if the current path matches.
var HighlightedLink = React.createClass({
    mixins: [Router.NavigatableMixin],

    isActive: function() {
        return this.getPath() === this.props.href;
    },

    render: function() {
        var className;

        if( this.props.activeClassName && this.isActive() ) {
            className = this.props.activeClassName;
        }

        var link = <Router.Link>{this.props.children}</Router.Link>;
        this.transferPropsTo(link);

        return <li className={className}>{link}</li>;
    }
});


var NavBar = React.createClass({
    render: function() {
        var navSpecs = [
            {path: '/',         icon: 'fa-dashboard',   text: 'Main Page'},
            {path: '/settings', icon: 'fa-cog',         text: 'Settings'},
            {path: '/about',    icon: 'fa-info-circle', text: 'About'},
        ];

        var navItems = map(navSpecs, function(item) {
            return (
                <HighlightedLink href={item.path} activeClassName="active">
                    <i className={'fa fa-fw ' + item.icon}></i> {item.text}
                </HighlightedLink>
            );
        });

        // TODO: make collapse work properly
        return (
            <nav className="navbar navbar-inverse navbar-fixed-top" role="navigation">
                <div className="navbar-header">
                    <button type="button" className="navbar-toggle">
                        <span className="sr-only">Toggle navigation</span>
                        <span className="fa fa-fw fa-inverse fa-bars"></span>
                    </button>
                    <a className="navbar-brand" href="/">Dashboard</a>
                </div>

                <div className="collapse navbar-collapse">
                    <ul className="nav navbar-nav side-nav">
                        {navItems}
                    </ul>
                </div>
            </nav>
        );
    },
});


module.exports = NavBar;
