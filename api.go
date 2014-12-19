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

// NewAPI returns a API struct
func NewAPI() *API {
	api := &API{
		HTTPClient: http.DefaultClient,
	}

	return api
}

func (a *API) apiGet(url string, v url.Values, data respObj) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = v.Encode()

	resp, err := a.HTTPClient.Do(req)
	defer resp.Body.Close()

	return decode(resp, data)
}

func (a *API) apiPost(url string, v url.Values, data respObj) error {
	resp, err := a.HTTPClient.PostForm(url, v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decode(resp, data)
}

func decode(resp *http.Response, data respObj) error {
	err := json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return err
	}
	return data.getError()
}
