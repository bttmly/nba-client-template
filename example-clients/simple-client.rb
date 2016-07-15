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

module NbaClient
  
  DATA = JSON.parse(File.read(File.expand_path("../../nba.json", __FILE__)))

  class Endpoint
    
    @@parameters = {}
    DATA["parameters"].each do |item|
      @@parameters[item["name"]] = item
    end

    attr_accessor :name, :defaults, :url

    def initialize(info)
      @parameters = info["parameters"]
      @name = snake(info["name"])
      @url = info["url"]
      @defaults = @parameters.inject({}) do |defaults, param|
        defaults[param] = @@parameters[param]["default"]
        defaults
      end
    end
  end

  class QueryString
    def initialize(params)
      query = []
      params.each {|param, value| query << URI.encode("#{param}=#{value}") }
      @query_str = "?" + query.join("&")
    end

    def to_s
      @query_str
    end
  end

  DATA["stats_endpoints"].each do |ep|
    endpoint = Endpoint.new ep
    define_method(endpoint.name) do |args|
      params = endpoint.defaults.merge args
      query_str = QueryString.new(params).to_s
      uri = URI(endpoint.url + query_str)
      JSON.parse open(uri, { "User-Agent" => DATA["user_agent"], "Referrer" => DATA["referrer"] }).read
    end
  end

  extend self
end

# usage:
puts NbaClient.player_profile({ "PlayerID" => 201939 })

