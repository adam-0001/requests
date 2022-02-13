package requests

import (
	"strings"
	"testing"

	"github.com/adam-0001/requests/functions"
)

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
	_, j, _ := functions.MakeBodyFromData(map[string]string{"hello": "world"})
	if j != "application/json" {
		t.Errorf("incorrect content-Type. Wanted: application/json, got: %s", j)
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
	r, err := s.MakeRequest("GET", "https://httpbin.org/get", nil, headers, nil, nil, false)
	if err != nil {
		t.Error(err)
	}
	// ERROR HAS TO DO WITH PSUEDO HEADER ORDER BEING NULL : SEE HEADERS.GO
	// fmt.Println(r.Text())
	if r.StatusCode() != 200 {
		t.Errorf("incorrect status code. Wanted: 200, got: %d", r.StatusCode())
	}
}

func TestRequestWithData(t *testing.T) {
	// _, j, _ := functions.MakeBodyFromData(test{"meow", 42})
	// if j != "application/json" {
	// 	t.Errorf("incorrect content-Type. Wanted: application/json, got: %s", j)
	// }
	_, j, _ := functions.MakeBodyFromData(map[string]string{"hello": "world"})
	if j != "application/json" {
		t.Errorf("incorrect content-Type. Wanted: application/json, got: %s", j)
	}

	s, err := NewSession(20000, "")
	if err != nil {
		t.Error(err)
	}

	// jsonStr := []byte(`{"title":"Buy cheese and bread for breakfast."}`)

	resp, err := s.Post("https://httpbin.org/post", map[string]string{"hello": "world"}, nil, nil, "flop=itsworking", true)
	if err != nil {
		t.Error(err)
	}
	resp1, err := s.Post("https://httpbin.org/post", nil, nil, nil, map[string]string{"hello": "world"}, true)
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
	if !strings.Contains(text, "hello") {
		t.Errorf("incorrect response. Wanted: hello in url params, got: %s", text)
	}

	text = resp1.Text
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(text, "world") {
		t.Errorf("incorrect response. Wanted: world in response body, got: %s", text)
	}
}
func TestDefer(t *testing.T) {
	s, err := NewSession(20000, "http://127.0.01:2311")
	if err != nil {
		t.Error(err)
	}
	headers := []map[string]string{
		{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"},
		{"Accept": "*/*"},
	}
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Recovered from panic %v", r)
		}
	}()
	r, err := s.MakeRequest("GET", "https://httpbin.org/get", nil, headers, nil, nil, false)
	if err != nil {
		t.Error(err)
	}
	if r.StatusCode() != 200 {
		t.Errorf("incorrect status code. Wanted: 200, got: %d", r.StatusCode())
	}
}
