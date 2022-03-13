package functions

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
	// if len(psuedoHeaderOrder) > 0 {
	// 	reqHeaders["PHeader-Order:"] = append(reqHeaders["PHeader-Order:"], psuedoHeaderOrder...)
	// }
	return reqHeaders
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func InferContentType(contentType string, headers *map[string][]string) {
	// 	_, ok := newHeaders["Content-Type"]
	// 	_, ok1 := newHeaders["content-type"]
	// 	if !ok && !ok1 {
	// 		newHeaders["Content-Type"] = []string{contentType}
	// 		newHeaders["Header-Order:"] = append(newHeaders["Header-Order:"], "content-type")
	// }

	//make a copy of the headers keys but in lowercase
	keysLower := make([]string, len(*headers))
	i := 0
	for k := range *headers {
		keysLower[i] = strings.ToLower(k)
		i++
	}

	//Check if keysLower contains "content-type"
	if !contains(keysLower, "content-type") {
		(*headers)["Content-Type"] = []string{contentType}
		(*headers)["Header-Order:"] = append((*headers)["Header-Order:"], "content-type")
	}

}

func FillNeededHeaders(host string, headers *http.Header) {
	defHeaders := defaultHeaders(host)
	keysLower := make([]string, len(*headers))
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
	// for k, v := range defHeaders {
	// 	key := strings.ToLower(k)

	// 	if _, ok := (*headers)[key]; !ok {
	// 		// fmt.Println("Edited:", key, ok)
	// 		(*headers)[k] = v
	// 		(*headers)["Header-Order:"] = append((*headers)["Header-Order:"], key)
	// 	}
	// }
}

func defaultHeaders(host string) map[string][]string {
	return (map[string][]string{
		"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36"},
		"Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"},
		"Accept-Encoding": {"gzip, deflate, br, utf-8"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"Connection":      {"keep-alive"},
		"Host":            {host},
	})
}
