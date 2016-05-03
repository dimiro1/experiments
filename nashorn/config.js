// Application default entry point and default protocol
var endpoint = 8080,
    protocol = "json";

/**
 * Implements the HelloService interface
 */
function HelloServiceHandler() {}

HelloServiceHandler.prototype.hello = function(name) {
    return "Hello " + name;
}

var helloService = new HelloServiceHandler();