package main

import (
	"fmt"

	"github.com/adam-0001/requests"
)

var headers = []map[string]string{
	{"hello": "world"},
	{"foo": "bar"},
}

func main() {
	resp, err := requests.Get("https://google.com", headers, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Text)
}
