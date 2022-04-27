package helpers

import (
	"strings"

	http "github.com/adam-0001/fhttp"
)

func MakeHeaders(headers []map[string]string) map[string][]string {
	var reqHeaders = map[string][]string{
		"Header-Order:":  {},
		"PHeader-Order:": {":method", ":authority", ":scheme", ":path"},
	}
	for _, v := range headers {

		for key, val := range v {
			reqHeaders[key] = []string{val}
			reqHeaders["Header-Order:"] = append(reqHeaders["Header-Order:"], strings.ToLower(key))
		}
	}

	return reqHeaders
}

func contains[T comparable](s []T, str T) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func InferContentType(contentType string, headers *map[string][]string) {
	//make a copy of the headers keys but in lowercase
	keysLower := make([]string, len(*headers))
	i := 0
	for k := range *headers {
		keysLower[i] = strings.ToLower(k)
		i++
	}

	if !contains(keysLower, "content-type") {
		(*headers)["Content-Type"] = []string{contentType}
		(*headers)["Header-Order:"] = append((*headers)["Header-Order:"], "content-type")
	}

}

func FillNeededHeaders(host string, headers *http.Header) {
	defHeaders := defaultHeaders(host)
	keysLower := make([]string, len(defHeaders))
	i := 0
	for k := range *headers {
		keysLower[i] = strings.ToLower(k)
		i++
	}
	for k := range defHeaders {
		lower := strings.ToLower(k)
		if !contains(keysLower, lower) {
			(*headers)[k] = defHeaders[k]
			(*headers)["Header-Order:"] = append((*headers)["Header-Order:"], lower)
		}
	}
}

func defaultHeaders(host string) map[string][]string {
	return (map[string][]string{
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36"},
		"Accept":     {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
		// "Accept-Encoding": {"gzip, deflate, br"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"Connection":      {"keep-alive"},
		"Host":            {host},
	})
}
