package requests

import (
	"strings"
	"testing"

	functions "github.com/adam-0001/requests/helpers"
)

func TestQuery(t *testing.T) {
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
	newHeaders := functions.MakeHeaders(headers)

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

func TestRequest(t *testing.T) {
	s, err := NewSession(20000, "")
	if err != nil {
		t.Error(err)
	}
	headers := []map[string]string{
		{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"},
		{"Accept": "*/*"},
	}
	r, err := s.MakeRequest("GET", "https://httpbin.org/get", headers, nil)
	if err != nil {
		t.Error(err)
	}

	if r.StatusCode() != 200 {
		t.Errorf("incorrect status code. Wanted: 200, got: %d", r.StatusCode())
	}
}

func TestRequestWithData(t *testing.T) {

	s, err := NewSession(20000, "")
	if err != nil {
		t.Error(err)
	}

	resp, err := s.Post("https://httpbin.org/post", nil, "flop=itsworking")
	if err != nil {
		t.Error(err)
	}
	resp1, err := s.Post("https://httpbin.org/post", nil, map[string]string{"hello": "world"})
	if err != nil {
		t.Error(err)
	}
	text := resp.Text
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(text, "flop") {
		t.Errorf("incorrect response. Wanted: flop in response body, got: %s", text)
	}

	text = resp1.Text
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(text, "world") {
		t.Errorf("incorrect response. Wanted: world in response body, got: %s", text)
	}
}
