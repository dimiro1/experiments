def sort(arr)
    return arr if arr.size <= 1
    
    pivot = arr.sample
    return sort(arr.select { |n| n < pivot }) + 
           arr.select { |n| n == pivot } + 
           sort(arr.select { |n| n > pivot })
end

arr = [10, 9, 2, 8, 7, 6, 1, 5, 4, 3, 2, 1]

puts sort(arr).to_s