/** @jsx React.DOM */

var React = require('react');

var NavBar = require('../components/NavBar.jsx');


var AboutPage = React.createClass({
    render: function() {
		return (
            <div id='wrapper'>
                <NavBar />
				<div id="page-wrapper">
					<div className="container-fluid">
						<div className="row">
							<div className="col-lg-12">
								<h1 className="page-header">
									About
								</h1>
							</div>
						</div>

						<div className="row">
							<div className="col-lg-12">
								This is a simple self-hosted dashboard that can display various types
								of data.  For more information, or to report an issue, please
								see <a href="https://github.com/andrew-d/dashboard">the GitHub page</a>.
							</div>
						</div>
					</div>
				</div>
			</div>
		);
	}
});

module.exports = AboutPage;
