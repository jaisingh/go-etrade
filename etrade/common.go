//Package etrade wraps the etrade rest api.
package etrade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Common functions

// unMarshalResponse is a wrapper for json.Unmarshal
func unmarshalResponse(resp *http.Response, out interface{}) error {
	data, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(data, out)

	return err
}

// Get wraps around the oauth.Get function and passes in
// the credentails from the OauthConfig and returns http.Response and error
func (c *OauthClient) Get(url string, opts map[string]string) (resp *http.Response, err error) {
	log.Println("Getting ", url)
	return c.Consumer.Get(url, opts, &c.Config.AccessToken)
}

// GetUnmarshall takes a url, options and the target object
// retrieves the data from etrade and unmarshals the contents into the object
func (c *OauthClient) GetUnmarshal(url string, opts map[string]string, v interface{}) error {
	resp, err := c.Get(url, opts)
	if err != nil {
		resp.Body.Close()
		return err
	}
	defer resp.Body.Close()

	err = unmarshalResponse(resp, v)
	if err != nil {
		return err
	}

	return nil
}

// DateTime field to handle etrade formatting
type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	ts, err := time.Parse(DATETIME_FORMAT, fmt.Sprintf("%s", b))
	t.Time = ts
	return err
}
