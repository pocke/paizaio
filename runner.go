package paizaio

import "net/url"

type Runner struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

type Details struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Note     string `json:"note"`
	Status   string `json:"status"`

	BuildStdout   string `json:"build_stdout"`
	BuildStderr   string `json:"build_stderr"`
	BuildExitCode int    `json:"build_exit_code"`
	BuildTime     string `json:"build_time"`
	BuildMemory   int    `json:"build_memory"`
	BuildResult   string `json:"build_result"`

	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exit"`
	Time     string `json:"time"`
	Memory   int    `json:"memory"`
	Result   string `json:"result"`
}

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
