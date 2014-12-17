package paizaio

import "net/http"

const (
	BaseURL = "http://api.paiza.io"
)

type API struct {
	HttpClient *http.Client
}
