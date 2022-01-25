package main

import (
	"fmt"
	"requests"
)

func main() {
	s, err := requests.NewSession(20000, "")
	if err != nil {
		panic(err)
	}
	headers := []map[string]string{
		{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"},
		{"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
	}
	r, err := s.MakeRequest("GET", "http://httpbin.org/get", nil, headers, nil, nil, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.StatusCode())
	// fmt.Println(r.Headers())
	resp, err := r.Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
