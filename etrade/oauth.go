package etrade

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/jaisingh/oauth"
)

// Oauth Client code

func NewOauthClient(config OauthConfig, opts ...string) *OauthClient {
	c := newOauthClient(config.ConsumerKey, config.ConsumerSecret, opts[0])
	c.Config = config

	return c
}

func newOauthClient(key string, secret string, opts ...string) *OauthClient {
	c := oauth.NewConsumer(
		key,
		secret,
		oauth.ServiceProvider{
			RequestTokenUrl:   REQUEST_TOKEN_URL,
			AuthorizeTokenUrl: AUTHORIZE_URL,
			AccessTokenUrl:    ACCESS_TOKEN_URL,
		})
	c.AdditionalAuthorizationUrlParams = map[string]string{
		"key": key,
	}

	if opts != nil && len(opts) > 0 {
		proxyUrl, _ := url.Parse(opts[0])
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(proxyUrl),
		}
		newHttpClient := &http.Client{Transport: tr}
		p := reflect.ValueOf(c)
		httpClient := p.Elem().FieldByName("HttpClient")
		httpClient.Set(reflect.ValueOf(newHttpClient))
	}

	return &OauthClient{*c, OauthConfig{
		ConsumerKey:    key,
		ConsumerSecret: secret,
	}}

}

type OauthClient struct {
	oauth.Consumer
	Config OauthConfig
}

type OauthConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    oauth.AccessToken
	TimeStamp      time.Time
}
