package paizaio

import "errors"
import "fmt"
import "net/url"

func (a *API) RunnerCreate(v url.Values) (*Runner, error) {
	lang := v.Get("language")
	included := false
	for _, v := range Languages {
		if v == lang {
			included = true
			break
		}
	}
	if !included {
		return nil, errors.New(fmt.Sprint("%s is not allowed language.", lang))
	}

	resp := &Runner{}
	err := a.apiPost(BaseURL+"/runners/create", v, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *API) RunnerStatus(id string) (*Runner, error) {
	v := url.Values{}
	v.Add("id", id)

	resp := &Runner{}
	err := a.apiGet(BaseURL+"/runners/get_status", v, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *API) RunnerDetails(id string) (*Details, error) {
	v := url.Values{}
	v.Add("id", id)

	resp := &Details{}
	err := a.apiGet(BaseURL+"/runners/get_details", v, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
