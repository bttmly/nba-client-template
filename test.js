var assert = require("assert");

// ensure JSON is valid
var template = require("./nba.json");

function assertIsString (s) {
  if (typeof s !== "string") {
    throw new Error(s + " is not a string")
  }
}

assertIsString(template.user_agent);
assertIsString(template.referrer);

template.stats_endpoints.forEach(function (endpoint) {
  // console.log(endpoint.name);
  assertIsString(endpoint.name);
  assertIsString(endpoint.url);
  endpoint.parameters.forEach(assertIsString);
  if (Object.keys(endpoint).length !== 3) {
    throw new Error(endpoint.name + " has more than three properties")
  }
});

template.parameters.forEach(function (parameter) {
  // console.log(parameter.name);
  assertIsString(parameter.name);
  assertIsString(parameter.default);
  parameter.values.forEach(assertIsString);
  if (Object.keys(parameter).length !== 3) {
    throw new Error(parameter.name + " has more than three properties")
  }
});