/** @jsx React.DOM */

var React = require('react'),
    RRouter = require('rrouter');

var Link = RRouter.Link;


var NavBar = React.createClass({
    render: function() {
        // TODO: set active
        var navItems = [];

        navItems.push(
            <li key="index" className="active">
                <Link to="index"><i className="fa fa-fw fa-dashboard"></i> Main Page</Link>
            </li>
        );
        navItems.push(
            <li key="settings">
                <Link to="settings"><i className="fa fa-fw fa-cog"></i> Settings</Link>
            </li>
        );

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
