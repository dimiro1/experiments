local helpers = require("helpers")
 
enpoint = ":8080"
protocol = "json"

service = {
    ret = "Hello World"
}

helpers.registerService("/service/", service)

print(offer)
print(offer.id)
print(offer["id"])

function hello(name)
    if name == "Claudemiro" then
        print("Welcome admin!")
        return true
    else
        print("Hello " .. name .. "!")
        return false
    end
end