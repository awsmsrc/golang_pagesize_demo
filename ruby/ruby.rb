require 'net/http'
require 'csv'


CSV.open("./output.csv", "w") do |output|
    CSV.foreach('../input.csv') do |row|
        size =  Net::HTTP.get_response(row[4], '/').body.length
        puts size
        output << (row << size)
    end
end
