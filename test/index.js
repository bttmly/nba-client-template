// ensure JSON is valid
const template = require("../nba.json");

function assertIsString (s) {
  if (typeof s !== "string") {
    throw new Error(s + " is not a string")
  }
}

assertIsString(template.user_agent);
assertIsString(template.referrer);

template.stats_endpoints.forEach((endpoint) => {
  // console.log(endpoint.name);
  assertIsString(endpoint.name);
  assertIsString(endpoint.url);
  if (endpoint.parameters.length === 0) {
    throw new Error(endpoint.name + " has no parameters")
  }
  endpoint.parameters.forEach(assertIsString);
  if (Object.keys(endpoint).length !== 3) {
    throw new Error(endpoint.name + " has more than three properties")
  }
});

template.parameters.forEach((parameter) => {
  // console.log(parameter.name);
  assertIsString(parameter.name);
  assertIsString(parameter.default);
  if (parameter.values.length === 0) {
    throw new Error(parameter.name + " has no values")
  }
  parameter.values.forEach(assertIsString);
  if (Object.keys(parameter).length !== 3) {
    throw new Error(parameter.name + " has more than three properties")
  }
});
