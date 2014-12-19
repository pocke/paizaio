package paizaio_test

import (
	"net/url"
	"testing"

	"github.com/pocke/paizaio"
)

var api = paizaio.NewAPI()
var idCh = make(chan string, 1)
var getID = func() func() string {
	var id string
	return func() string {
		if id != "" {
			return id
		}
		id = <-idCh
		return id
	}
}()

const code = `
package main
import "fmt"

func main() {
  fmt.Println("Hello World!")
}
`

func TestRunnerCreate(t *testing.T) {
	v := url.Values{}
	v.Add("language", "go")
	v.Add("source_code", code)

	r, err := api.RunnerCreate(v)
	t.Logf("%+v", r)
	if err != nil {
		t.Error(err)
	}
	idCh <- r.ID

	v.Set("language", "Golang")
	_, err = api.RunnerCreate(v)
	if err == nil {
		t.Error("Should raise error when dont allowed language. But not raise.")
	}
}

func TestRunnerStatus(t *testing.T) {
	id := getID()
	r, err := api.RunnerStatus(id)
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestRunnerDetails(t *testing.T) {
	id := getID()
	d, err := api.RunnerDetails(id)
	if err != nil {
		t.Error(err)
	}
	t.Log(d)
}
