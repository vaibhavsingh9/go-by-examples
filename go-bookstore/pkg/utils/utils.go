package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// to unmarshal the request we receive in json
func ParseBody(r *http.Request, X interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal([]byte(body), X); err != nil {
			return
		}
	}
}
