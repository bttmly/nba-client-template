# run this script with the python executable to see output

import os
import json
import urllib2
import re

# wtf python?
# http://stackoverflow.com/a/13105359
def byteify(input):
    if isinstance(input, dict):
        return {byteify(key): byteify(value)
                for key, value in input.iteritems()}
    elif isinstance(input, list):
        return [byteify(element) for element in input]
    elif isinstance(input, unicode):
        return input.encode('utf-8')
    else:
        return input

def snake(name):
  s1 = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', name)
  return re.sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()

### real code starts here...

# acquire data from json files
path = os.path.join(os.path.dirname(__file__), '../nba.json')
data = byteify(json.loads(open(path).read()))

parameters = {}
for item in data["parameters"]:
  parameters[item["name"]] = item

# use this basically as just a name space to put our api call methods on
class NbaClient(object):
  pass

# make a function that gets data from a specific end point with built-in default parameter values
def make_endpoint(name, endpoint):
  def func(params):
    all_params = dict()

    for param in endpoint["parameters"]:
      value = params.get(param, parameters[param]["default"])
      all_params[param] = value

    query = []
    for param, value in all_params.iteritems():
      query.append(urllib2.quote(param) + '=' + urllib2.quote(str(value)))
    query_str = '?' + '&'.join(query)

    url = endpoint['url'] + query_str
    req = urllib2.Request(url)
    req.add_header('Referrer', data['referrer'])
    req.add_header('User-Agent', data['user_agent'])
    return byteify(json.loads(urllib2.urlopen(req).read()))

  func.__name__ = name
  return func

# iterate through the endpoint data and add a static method for each one
for endpoint in data['stats_endpoints']:
  func = make_endpoint(endpoint["name"], endpoint)
  setattr(NbaClient, snake(endpoint["name"]), staticmethod(func))

# usage:
print NbaClient.player_profile({ 'PlayerID': 201939 })
