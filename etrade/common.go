package etrade

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Common functions

func unmarshalResponse(resp *http.Response, out interface{}) error {
	data, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(data, out)

	return err
}
