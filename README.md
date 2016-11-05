# nba-client-template

This repo holds a JSON file that can be used to programmatically generate a client for the NBA's API. Since it's JSON, it should be trivially parseable by any language. It's currently published as an NPM module, but if you want it packaged up for your language of choice as well, please open an [issue](https://github.com/nickb1080/nba-client-template/issues) or better yet, a [pull request](https://github.com/nickb1080/nba-client-template/pulls). Thanks!

The file contains two primary parts: endpoints and parameters. Endpoints generally should be mapped to methods on an API client. Each endpoint has a list of parameters that it accepts. Parameter objects hold some data about individual parameters. Specifically, what the parameter is called (`name`), it's known values (`values`), it's default value (`default`), and a boolean indicating whether it's listed values encompass all valid values (`enumerated`). These can be used to implement parameter validation.

The simplest and probably easiest implementation is to map each endpoint to a function accepting a single argument, that being a hash of `{parameter: value}`. To provide some examples in code, the `examples` folder contains three very, very simple client implementations, in Python, Ruby, and Go. The Python and Ruby examples are implemented pretty much the same way, and are certainly not idiomatic in the language, but should illustrate the general idea. For languages without highly dynamic reflection features, code generation might be used to create a client. The Go example illustrates how that might work -- using Go to generate a Go file.

A full-featured Node.js client based on this template can be seen at [nickb1080/nba](https://github.com/nickb1080/nba).
