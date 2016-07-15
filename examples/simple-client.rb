require "json"
require "uri"
require "open-uri"

# this client doesn't implement parameter validation, though it can be derived
# from the JSON data as well.

# holds 'static' API call methods
module NbaClient
  data = JSON.parse(File.read(File.expand_path("../../endpoints.json", __FILE__))) 

  data["endpoints"].each do |name, endpoint|
    define_method(name) do |params|
      all_params = endpoint["defaults"].merge params
      
      query = []
      all_params.each {|param, value| query << URI.encode("#{param}=#{value}") }
      query_str = "?" + query.join("&")
      
      url = endpoint["url"] + query_str
      uri = URI(url)
      
      JSON.parse uri.read
    end
  end

  extend self
end

# usage:
puts NbaClient.player_profile({ "PlayerID" => 201939 })

