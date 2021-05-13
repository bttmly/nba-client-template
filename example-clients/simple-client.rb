require "json"
require "uri"
require "open-uri"
require 'net/http'

def snake(str)
  str
    .gsub(/::/, '/')
    .gsub(/([A-Z]+)([A-Z][a-z])/,'\1_\2')
    .gsub(/([a-z\d])([A-Z])/,'\1_\2')
    .tr("-", "_")
    .downcase
end

def result_set_to_hash (result_set)
  headers = result_set[:headers]
  values = result_set[:rowSet]

end

# this client doesn't implement parameter validation, though it can be derived
# from the JSON data as well.

class NbaClient

  TEMPLATE = JSON.parse(File.read(File.expand_path("../../nba.json", __FILE__)),
    symbolize_names: true
  )

  PARAMETERS = TEMPLATE[:parameters].each_with_object({}) do |param, obj|
    name = param[:name].to_sym
    obj[name] = param
  end

  HEADERS = {
    "Accept-Language" => "en-US",
    "Accept" => "*/*",
    "User-Agent" => TEMPLATE[:user_agent],
    "Referer" => TEMPLATE[:referrer],
    "Connection" => "keep-alive",
    "Cache-Control" => "no-cache",
    "Origin" => "http://stats.nba.com",
  }.freeze

  class Endpoint
    attr_reader :defaults, :name, :url
    def initialize(info)
      @parameters = info[:parameters]
      @name = snake(info[:name])
      @url = info[:url]
      @defaults = @parameters.each_with_object({}) do |name, defaults|
        name = name.to_sym
        defaults[name] = PARAMETERS[name][:default]
      end.freeze
    end

    def fetch (params_in, extra_headers, host_override = nil)
      params = { **@defaults, **params_in}
      url = URI("#{@url}?#{URI.encode_www_form(params)}")
      url.host = host_override unless host_override.nil?
      Net::HTTP.get_response(url, { **HEADERS, **extra_headers })
    end
  end

  def initialize (headers: nil, host: nil)
    @headers = headers || {}
    @host = host
  end

  TEMPLATE[:stats_endpoints].each do |endpoint_config|
    endpoint = Endpoint.new(endpoint_config)
    define_method(endpoint.name) do |params|
      res = endpoint.fetch(params, @headers, @host)
      status = res.code.to_i
      if status >= 400
        raise "HTTP error #{status}: #{res.body}"
      end
      JSON.parse(res.body, symbolize_names: true)
    end
  end
end

# usage:
nba = NbaClient.new
steph = nba.player_info({ PlayerID: 201939 })

wnba = NbaClient.new(host: "stats.wnba.com")
wnba.player_info({ PlayerID: 1628886 })

# helper method for transforming responses
def response_to_hash (res_data)
  res_data[:resultSets].each_with_object({}) do |rs, obj|
    keys = rs[:headers]
    values = rs[:rowSet].fetch(0, [])
    obj[rs[:name]] = Hash[keys.zip(values)]
  end
end

puts response_to_hash(steph)