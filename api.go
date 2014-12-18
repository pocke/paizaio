package paizaio

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	BaseURL = "http://api.paiza.io"
)

type API struct {
	HTTPClient *http.Client
}

func NewAPI() *API {
	api := &API{
		HTTPClient: http.DefaultClient,
	}

	return api
}

func (a *API) apiGet(url string, v url.Values, data interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = v.Encode()

	resp, err := a.HTTPClient.Do(req)
	defer resp.Body.Close()

	return decode(resp, data)
}

func (a *API) apiPost(url string, v url.Values, data interface{}) error {
	resp, err := a.HTTPClient.PostForm(url, v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decode(resp, data)
}

func decode(resp *http.Response, data interface{}) error {
	return json.NewDecoder(resp.Body).Decode(data)
}
