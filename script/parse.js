const fs = require("fs");
const path = require("path");
const url = require("url");

const template = require("../nba");
const nameIndexedParams = {};
template.parameters.forEach(p => nameIndexedParams[p.name] = p);

const urls = fs.readFileSync(path.join(__dirname, "/urls"), "utf8")
  .split(/\n/)
  .filter(Boolean)
  .map(uriToEndpoint)
  .map(verifyParametersExist)
  .map(ep => JSON.stringify(ep, null, 2))
  .map(ep => console.log(ep))


function uriToEndpoint (uri) {
  const parsed = url.parse(uri, true);
  console.log("parsing", parsed.pathname);
  const parameters = Object.keys(parsed.query);
  return {
    url: `${parsed.protocol}//${parsed.host}${parsed.pathname}`,
    parameters,
  }
}

function verifyParametersExist (endpoint) {
  endpoint.parameters.forEach(function (p) {
    if (nameIndexedParams[p] == null) {
      console.error("Unknown parameter", p, endpoint.url);
    }
  });
  return endpoint;
}