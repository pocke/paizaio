package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pocke/paizaio"
)

const code = `
package main
import "fmt"

func main() {
  fmt.Println("Hello World!")
}
`

func main() {
	paiza := paizaio.NewAPI()
	v := url.Values{}
	v.Add("language", "go")
	v.Add("source_code", code)

	r, err := paiza.RunnerCreate(v)
	fmt.Printf("%+v\n", r)
	if err != nil {
		panic(err)
	}

	for {
		s, err := paiza.RunnerStatus(r.ID)
		if err != nil {
			panic(err)
		}
		if s.Status != paizaio.StatusRunning {
			break
		}

		time.Sleep(10 * time.Second)
	}
	d, err := paiza.RunnerDetails(r.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", d)
}
