require "json"
require "uri"
require "open-uri"

def snake(str)
  str.gsub(/::/, '/').
  gsub(/([A-Z]+)([A-Z][a-z])/,'\1_\2').
  gsub(/([a-z\d])([A-Z])/,'\1_\2').
  tr("-", "_").
  downcase
end

# this client doesn't implement parameter validation, though it can be derived
# from the JSON data as well.



# holds 'static' API call methods
module NbaClient
  data = JSON.parse(File.read(File.expand_path("../../nba.json", __FILE__))) 

  parameters = {}
  data["parameters"].each do |item|
    parameters[item["name"]] = item
  end

  data["stats_endpoints"].each do |endpoint|
    name = snake(endpoint["name"])

    define_method(name) do |params|
      all_params = {}
      endpoint["parameters"].each do |param|
        value = params.fetch param, parameters[param]["default"]
        all_params[param] = value || default
      end
      
      query = []
      all_params.each {|param, value| query << URI.encode("#{param}=#{value}") }
      query_str = "?" + query.join("&")
      
      url = endpoint["url"] + query_str
      uri = URI(url)
      JSON.parse open(uri, { "User-Agent" => data["user_agent"], "Referrer" => data["referrer"] }).read
    end
  end

  extend self
end

# usage:
puts NbaClient.player_profile({ "PlayerID" => 201939 })

