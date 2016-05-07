local helpers = require("helpers")
 
enpoint = ":8080"
protocol = "json"

service = {
    ret = "Hello World"
}

helpers.registerService("/service/", service)