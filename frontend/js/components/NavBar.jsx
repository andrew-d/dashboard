/** @jsx React.DOM */

var React = require('react');


var NavBar = React.createClass({
    render: function() {
        var style = {
            //'background': 'rgb(37, 42, 58)',
            'border-right': '1px dashed gray',
            'height': '100%',
        };

        return (
            <div id="nav" className="col-sm-3" style={style}>
                <ul className="nav nav-stacked">
                    <li><a href="/">Main Page</a></li>
                </ul>
            </div>
        );
    },
});


module.exports = NavBar;
