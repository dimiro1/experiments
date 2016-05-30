local array = require "array"

a = array.new(1000)

print(#a)

for i=1,1000 do
    a[i] = 1/i
end

print(a[10])