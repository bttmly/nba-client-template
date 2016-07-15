# run this script with the python executable to see output

import os
import json
import urllib2

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

### real code starts here...

# acquire data from json files
path = os.path.join(os.path.dirname(__file__), "../endpoints.json")
data = byteify(json.loads(open(path).read()))

# use this basically as just a name space to put our api call methods on
class NbaClient(object):
  pass

# make a function that gets data from a specific end point with built-in default parameter values
def make_endpoint(name, endpoint):
  def func(params):
    all_params = dict()
    all_params.update(endpoint['defaults'])
    all_params.update(params)

    query = []
    for param, value in all_params.iteritems():
      query.append(urllib2.quote(param) + "=" + urllib2.quote(str(value)))
    query_str = "?" + "&".join(query)

    url = endpoint["url"] + query_str
    req = urllib2.Request(url)
    req.add_header("Referrer", data["referrer"])
    req.add_header("User-Agent", data["user_agent"])
    return byteify(json.loads(urllib2.urlopen(req).read()))

  func.__name__ = name
  return func

# iterate through the endpoint data and add a static method for each one
for name, endpoint in data['endpoints'].iteritems():
  func = make_endpoint(name, endpoint)
  setattr(NbaClient, name, staticmethod(func))

# usage:
print NbaClient.player_profile({ "PlayerID": 201939 })
