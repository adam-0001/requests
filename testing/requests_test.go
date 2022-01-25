package main

import (
	"fmt"
	"requests"
	"requests/functions"
	"strings"
	"testing"
)

type test struct {
	helloworld  string `json:"helloworld"`
	mellowporld int    `json:"mellowporld"`
}

func TestQuery(t *testing.T) {
	// fmt.Println("Testing query")
	query := map[string]string{
		"q":     "test",
		"hello": "world",
	}
	url, _, err := functions.GetCompleteQuery("http://www.google.com", query)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(url, "hello=world") || !strings.Contains(url, "q=test") {
		t.Error("Wanted: q=test&hello=world, got:", url)
	}
}

func TestHeaders(t *testing.T) {
	// fmt.Println("Testing headers")
	headers := []map[string]string{
		{"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.1"},
		{"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
	}
	newHeaders := functions.MakeHeaders(headers, []string{":method", ":authority"})

	if newHeaders["Header-Order:"][0] != "user-agent" ||
		newHeaders["Header-Order:"][1] != "accept" {
		t.Errorf("Incorrect Header Order. Wanted: user-agent, accept, got: %s, %s", newHeaders["Header-Order:"][0], newHeaders["Header-Order:"][1])
	}
	if newHeaders["PHeader-Order:"][0] != ":method" ||
		newHeaders["PHeader-Order:"][1] != ":authority" {
		t.Error("Incorrect PHeader Order")
	}
	if newHeaders["User-Agent"][0] != "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.1" || newHeaders["Accept"][0] != "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8" {
		t.Error("Incorrect Header Values")
	}
}

func TestData(t *testing.T) {
	_, j, _ := functions.MakeBodyFromData(test{"meow", 42})
	if j != "application/json" {
		t.Errorf("incorrect content-Type. Wanted: application/json, got: %s", j)
	}
	_, j, _ = functions.MakeBodyFromData(map[string]string{"hello": "world"})
	if j != "application/json" {
		t.Errorf("incorrect content-Type. Wanted: application/json, got: %s", j)
	}
	// _, _, _ = functions.MakeBodyFromData([]byte("hello=world"))
	// fmt.Println(j)

}

func TestRequest(t *testing.T) {
	s, err := requests.NewSession(20000, "")
	if err != nil {
		t.Error(err)
	}
	headers := []map[string]string{
		{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"},
		{"Accept": "*/*"},
	}
	r, err := s.MakeRequest("GET", "https://httpbin.org/get", nil, headers, nil, nil, false)
	if err != nil {
		t.Error(err)
	}
	// ERROR HAS TO DO WITH PSUEDO HEADER ORDER BEING NULL : SEE HEADERS.GO
	fmt.Println(r.Text())
	if r.StatusCode() != 200 {
		t.Errorf("incorrect status code. Wanted: 200, got: %d", r.StatusCode())
	}
}
