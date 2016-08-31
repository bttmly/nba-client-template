
const template = require("../nba");
const request = require("request-promise");

const OVERRIDES = {
  PlayerID: 201939 // Steph Curry
};

const docs = []

const defaultsByName = {};
template.parameters.forEach(function (param) {
  defaultsByName[param.name] = param.default;
});

function getEndpoint (endpoint) {
  console.log(endpoint);

  const params = {};
  endpoint.parameters.forEach(function (param) {
    if (OVERRIDES[param]) {
      params[param] = OVERRIDES[param];
      return;
    }

    params[param] = defaultsByName[param];
  });


  console.log(params);

  request({
    url: endpoint.url,
    qs: params,
  }).then(data => 
    documentResponse(endpoint, JSON.parse(data)))
}

function documentResponse (endpoint, data) {
  Object.key(data).forEach(k => documentItem(k, data[k]));
}

function documentItem (key, data) {

}

function handleList (list) {

}

function handleStruct (struct) {

}

function handlePrimitive (key, val) {

}

function handlePair (key, value, depth) {
  if (isPrimitive(value)) {
    return `${key}: ${type(value)}`
  }

  if (isArray(value)) {
    return `${key}: []${handlePair(value[0])}`
  }

  return `${key}: $`
)
}

getEndpoint(template.stats_endpoints[0])
