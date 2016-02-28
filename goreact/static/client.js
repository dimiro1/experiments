var renderClient = function renderClient(name) {
    ReactDOM.render(
    	React.createElement(HelloMessage, { name: name }), document.getElementById("content")
    );
};