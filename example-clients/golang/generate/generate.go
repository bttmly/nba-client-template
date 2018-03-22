package main

// NOTE: run this from directory nba-client-template/example-clients/golang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Endpoint struct {
	Name       string   `json:"name"`
	URL        string   `json:"url"`
	Parameters []string `json:"parameters"`
}

type Parameter struct {
	Name         string   `json:"name"`
	DefaultValue string   `json:"default"`
	Values       []string `json:"values"`
}

type Template struct {
	UserAgent  string      `json:"user_agent"`
	Referrer   string      `json:"referrer"`
	Referer    string      `json:"referer"`
	Origin     string      `json:"origin"`
	Endpoints  []Endpoint  `json:"stats_endpoints"`
	Parameters []Parameter `json:"parameters"`
}

func upperFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

var functionTemplate string = "func %s(params map[string]string) (resp *http.Response, err error) {\n    return apiGet(\"%s\", params)\n}"

func makeMethodString(ep Endpoint) string {
	name := ""
	pieces := strings.Split(ep.Name, "_")
	for _, p := range pieces {
		name += upperFirst(p)
	}
	return fmt.Sprintf(functionTemplate, name, ep.URL)
}

func main() {
	file, e := ioutil.ReadFile("../../nba.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	tmpl := Template{}
	json.Unmarshal(file, &tmpl)
	// fmt.Printf("%+v", tmpl)

	var pieces []string
	pieces = append(pieces, "package nba")
	pieces = append(pieces, "// NOTE: This file was generated automatically, do not edit directly!")
	pieces = append(pieces, "import \"net/http\"")
	pieces = append(pieces, "const UserAgent = \""+tmpl.UserAgent+"\"")
	pieces = append(pieces, "const Referer = \""+tmpl.Referer+"\"")
	pieces = append(pieces, "const Origin = \""+tmpl.Origin+"\"")
	for _, ep := range tmpl.Endpoints {
		pieces = append(pieces, makeMethodString(ep))
	}

	bytes := []byte(strings.Join(pieces, "\n\n"))
	err := ioutil.WriteFile("./client/client.go", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
