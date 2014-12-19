package paizaio

type APIError struct {
	Body string
}

func newAPIError(body string) *APIError {
	return &APIError{Body: body}
}

func (a *APIError) Error() string {
	return a.Body
}

var _ error = &APIError{}
