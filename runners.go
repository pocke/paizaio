package paizaio

import "net/url"

func (a *API) RunnerCreate(v url.Values) (*Runner, error) {
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
