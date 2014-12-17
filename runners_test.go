package paizaio_test

import (
	"net/url"
	"testing"
)

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
}
