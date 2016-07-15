package nba

import (
  "net/http"
  "net/url"
  "bytes"
)

func mapToQueryString(data map[string]string) string { 
    var buf bytes.Buffer
    buf.WriteString("?")
    for k, v := range data {
        buf.WriteString(url.QueryEscape(k)) 
        buf.WriteString("=")
        buf.WriteString(url.QueryEscape(v)) 
        buf.WriteString("&")
    } 
    return buf.String()
} 

func apiGet (url string, params map[string]string) (*http.Response, error) {
  qs := mapToQueryString(params)
  return http.Get(url + qs)
}
