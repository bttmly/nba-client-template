fs = require "fs"
cc = require "change-case"

_stats_endpoints = require "../nba/lib/stats-endpoints.js"
_parameters = require "../nba/lib/params"

output =
  user_agent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36"
  referrer: "http://stats.nba.com/scores/"

stats_endpoints = []
for name, endpoint of _stats_endpoints
  stats_endpoints.push {
    name,
    url: endpoint.url
    parameters: endpoint.params or Object.keys endpoint.defaults
  }
output.stats_endpoints = stats_endpoints

parameters = []
for name, data of _parameters
  if name is "PlayType" then continue
  _default = data.__DEFAULT__
  delete data.__DEFAULT__

  values = Object.keys data

  if _default not in values
    values.push _default

  parameters.push {
    name
    default: _default
    values
  }

output.parameters = parameters

fs.writeFileSync "./nba.json", JSON.stringify output, null, 2