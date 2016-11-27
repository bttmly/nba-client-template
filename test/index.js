var assert = require("assert");
var qs = require("querystring");

// ensure JSON is valid
var template = require("../nba.json");

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
  if (endpoint.parameters.length === 0) {
    throw new Error(endpoint.name + " has no parameters")
  }
  endpoint.parameters.forEach(assertIsString);
  if (Object.keys(endpoint).length !== 3) {
    throw new Error(endpoint.name + " has more than three properties")
  }
});

template.parameters.forEach(function (parameter) {
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

// const _ = require("lodash");
// const nameIndexedParams = _.keyBy(template.parameters, "name");
// const async = require("async");
// const request = require("request");
//
// async.each(template.stats_endpoints, function (endpoint, next) {
//   const defaults = {};
//   endpoint.parameters.forEach(function (param) {
//     defaults[param] = nameIndexedParams[param].default;
//   });
//   const url = endpoint.url + "?" + qs.stringify(defaults);
//   request(url, function (err, resp) {
//     if (resp.statusCode !== 200) {
//       console.log("fail", endpoint.url, resp.body)
//     } else {
//       console.log("success", endpoint.url);
//     }
//     next(err)
//   })
// }, function (err) {
//   console.log("done:", err);
// });
