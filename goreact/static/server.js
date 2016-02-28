var renderServer = function (name) {
    return ReactDOMServer.renderToString(
        React.createElement(HelloMessage, { name: name })
    );
};