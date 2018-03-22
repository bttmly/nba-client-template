package nba

import (
	"net/http"
)

func apiGet(reqURL string, params map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL, nil)

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Referer", Referer)
	req.Header.Set("Origin", Origin)

	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
